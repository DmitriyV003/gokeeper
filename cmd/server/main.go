package main

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
	"gokeeper/internal/config"
	"gokeeper/internal/core/services"
	"gokeeper/internal/data/postgres"
	"gokeeper/internal/proto"
	"gokeeper/internal/server"
	services2 "gokeeper/internal/services"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := initConfig()
	pool := connectToDB(cfg.DBUri)
	if pool != nil {
		migrate(pool)
	}

	mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	srv2 := &http.Server{
		Addr:    cfg.GrpcServerPort,
		Handler: nil,
	}
	grpcServer := grpc.NewServer()
	var listen net.Listener
	var err error

	loginSecretRepo := postgres.NewLoginSecretRepository(pool)
	userRepo := postgres.NewUserRepository(pool)
	secretService := services.NewSecretService(loginSecretRepo)
	authService := services2.NewAuthService(userRepo, cfg.JWTSecret)
	keyService := services.NewKeysService(cfg.MasterPassword)
	userService := services2.NewUserService(keyService, userRepo)
	proto.RegisterLoginSecretServiceServer(grpcServer, server.NewLoginSecretServer(secretService))
	proto.RegisterUserServer(grpcServer, server.NewUserServer(authService, userService))

	g, gCtx := errgroup.WithContext(mainCtx)
	g.Go(func() error {
		listen, err = net.Listen("tcp", cfg.GrpcServerPort)
		return err
	})
	g.Go(func() error {
		<-gCtx.Done()
		grpcServer.Stop()
		listen.Close()
		return srv2.Shutdown(gCtx)
	})

	if err := g.Wait(); err != nil {
		log.Error().Err(err).Msg("Server down")
	}
	log.Info().Msg("Application shutdown")
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

type migration struct {
	ID   int64
	Name string
}
