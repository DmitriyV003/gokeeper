package commands

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gokeeper/internal/core/services"
)

type Deps struct {
	AuthService        *services.AuthService
	UserService        *services.UserService
	LoginSecretService *services.LoginSecretService
}

var (
	rootCmd = &cobra.Command{
		Use:   "gophkeeper",
		Short: "Passwords Manager gophkeeper",
	}

	deps Deps

	unauthorized = func(cmd *cobra.Command) {
		cmd.PrintErrln("You are not authorized")
	}
)

func Execute(ctx context.Context, depp Deps) {
	deps = depp
	rootCmd.SetContext(ctx)

	if err := rootCmd.Execute(); err != nil {
		log.Debug().Err(err).Msg("Root command execution")
		rootCmd.PrintErrln(err)
	}
}
