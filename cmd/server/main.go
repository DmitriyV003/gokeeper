package main

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
	"gokeeper/internal/config"
	"gokeeper/internal/core"
	"gokeeper/internal/data"
	"gokeeper/internal/proto"
	"gokeeper/internal/server"
	"gokeeper/internal/server/handlers"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	_ = initConfig()
	//pool := connectToDB(cfg.DBUri)
	pool := connectToDB("postgres://homestead:homestead@localhost:54321/homestead")
	if pool != nil {
		migrate(pool)
	}

	mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: createHandler(pool),
		BaseContext: func(listener net.Listener) context.Context {
			return mainCtx
		},
	}

	srv2 := &http.Server{
		Addr:    ":8082",
		Handler: nil,
	}
	var grpcServer *grpc.Server
	var listen net.Listener
	var err error
	proto.RegisterLoginServiceServer(grpcServer, server.NewLoginSecretServer())

	g, gCtx := errgroup.WithContext(mainCtx)
	g.Go(func() error {
		grpcServer = grpc.NewServer()
		listen, err = net.Listen("tcp", ":8082")
		return err
	})
	g.Go(func() error {
		return srv.ListenAndServe()
	})
	g.Go(func() error {
		return srv2.ListenAndServe()
	})
	g.Go(func() error {
		<-gCtx.Done()
		log.Warn().Msg("Server down")
		grpcServer.Stop()
		listen.Close()
		srv2.Shutdown(gCtx)
		return srv.Shutdown(gCtx)
	})

	if err := g.Wait(); err != nil {
		log.Error().Err(err).Msg("Server down")
	}
}

func initConfig() config.Config {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msgf("can not load config")
	}

	return *cfg
}

func connectToDB(dbUri string) (pool *pgxpool.Pool) {
	if dbUri == "" {
		log.Warn().Msg("Database URl not provided")
		return nil
	}

	var err error
	conf, err := pgxpool.ParseConfig(dbUri)
	if err != nil {
		log.Error().Err(err).Msg("Unable to parse Database config")
		return
	}
	pool, err = pgxpool.ConnectConfig(context.Background(), conf)

	if err != nil {
		log.Error().Err(err).Msg("Unable to connect to database")
		return
	}

	return pool
}

func migrate(pool *pgxpool.Pool) {
	createMigrationsTable(pool)
	log.Info().Msgf("Creating migrations table")

	migrations, err := os.ReadDir("migrations")
	if err != nil {
		log.Error().Err(err).Msg("unable to read migrations directory")
		return
	}

	var file *os.File

	for _, migrationFile := range migrations {
		file, err = os.Open(fmt.Sprintf("migrations/%s", migrationFile.Name()))
		if err != nil {
			log.Error().Err(err).Msgf("unable to open migrationFile: %s", migrationFile.Name())
			return
		}

		sql := `SELECT id, name FROM migrations WHERE name = $1`
		var dbMigration migration

		err = pool.QueryRow(context.Background(), sql, migrationFile.Name()).Scan(&dbMigration.ID, &dbMigration.Name)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			log.Error().Err(err).Msg("Error to query migration")
			return
		}

		if err == nil {
			log.Info().Msgf("Migration passed: %s", migrationFile.Name())
			continue
		}

		wr := bytes.Buffer{}
		sc := bufio.NewScanner(file)
		for sc.Scan() {
			text := sc.Text()
			if text == "---- create above / drop below ----" {
				break
			}
			wr.WriteString(sc.Text())
		}

		sql = `INSERT INTO migrations (name) VALUES ($1)`
		_, err = pool.Exec(context.Background(), sql, migrationFile.Name())
		if err != nil {
			log.Error().Err(err).Msgf("unable to write migration to database: %s", migrationFile.Name())
			return
		}

		_, err = pool.Exec(context.Background(), wr.String())
		if err != nil {
			log.Error().Err(err).Msg("Error during migrationFile")
			return
		}

		log.Info().Msgf("Migrating: %s", migrationFile.Name())
	}
	log.Info().Msg("Migrations passed successfully")
}

func createMigrationsTable(pool *pgxpool.Pool) {
	sql := `CREATE TABLE IF NOT EXISTS migrations(
    	id serial PRIMARY KEY,
    	name VARCHAR (255) NOT NULL UNIQUE)`
	_, err := pool.Exec(context.Background(), sql)
	if err != nil {
		log.Error().Err(err).Msg("Error during migrationFile")
		return
	}
}

func createHandler(pool *pgxpool.Pool) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.StripSlashes)
	router.Use(middleware.Compress(5))
	router.Use(middleware.Heartbeat("/heartbeat"))

	userRepository := data.NewUserRepository(pool)
	userService := core.NewAuthService("SECRET", userRepository)

	router.Route("/api", func(r chi.Router) {
		r.Post("/register", handlers.NewRegisterHandler(userService).Handle())
		r.Post("/login", handlers.NewLoginHandler(userService).Handle())
	})

	return router
}

type migration struct {
	ID   int64
	Name string
}
