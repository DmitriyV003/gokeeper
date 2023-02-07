package main

import (
	"context"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gokeeper/cmd/cli/commands"
	"gokeeper/internal/client"
	config2 "gokeeper/internal/config"
	services2 "gokeeper/internal/core/services"
	sqlite2 "gokeeper/internal/data/sqlite"
	"os"
)

const notAssigned = "N/A"

var (
	buildVersion = notAssigned
	buildTime    = notAssigned
	buildCommit  = notAssigned
)

func main() {
	cfg := initConfig()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "03:04:05PM"})

	ctx := context.Background()

	log.Info().Msg(fmt.Sprintf("Build version: %s", buildVersion))
	log.Info().Msg(fmt.Sprintf("Build date: %s", buildTime))
	log.Info().Msg(fmt.Sprintf("Build commit: %s\n", buildCommit))

	db, err := sqlite2.NewSQLite(ctx, cfg.SQLiteUri)
	if err != nil {
		log.Fatal().Err(err).Msg("Connecting to the SQLite database")
	}

	settingsRepo := sqlite2.NewSettingsRepository(db)
	authService := services2.NewAuthService(cfg.JWTSecret, settingsRepo)
	keysService := services2.NewKeysService("")
	deps := commands.Deps{
		AuthService:        authService,
		UserService:        services2.NewUserService(client.NewUserClient(ctx, ":8082"), settingsRepo),
		LoginSecretService: services2.NewLoginSecretService(authService, client.NewLoginSecretClient(ctx, ":8082"), settingsRepo, "", keysService),
	}

	commands.Execute(ctx, deps)
}

func initConfig() config2.Config {
	cfg, err := config2.Load()
	if err != nil {
		log.Fatal().Err(err).Msgf("can not load config")
	}

	return *cfg
}
