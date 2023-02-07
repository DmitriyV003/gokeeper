package commands

import (
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gokeeper/internal/data"
)

func init() {
	rootCmd.AddCommand(registerCmd)
}

var registerCmd = &cobra.Command{
	Use:       "register",
	Short:     "registers you in the system",
	Args:      cobra.MinimumNArgs(2),
	ValidArgs: []string{"login", "password"},
	Run: func(cmd *cobra.Command, args []string) {
		authorized, err := deps.AuthService.CheckAuthorized(cmd.Context())
		if authorized || errors.Is(err, data.ErrLoggedInAlready) {
			cmd.PrintErrln("You are already logged into the system. " +
				"If you want to register a new user then logout first.")
			return
		}
		if err != nil {
			log.Error().Err(err).Msg("Checking whether authorized user")
			return
		}

		login, password := args[0], args[1]

		if err := deps.UserService.Register(cmd.Context(), login, password); err != nil {
			if errors.Is(err, data.ErrLoginTaken) {
				cmd.PrintErrln("This login is already taken. Try another one.")
				return
			}

			log.Error().Err(err).Msg("Registering user from command")
			return
		}

		cmd.Printf("Registered [%s] successfully!\n", login)
	},
}
