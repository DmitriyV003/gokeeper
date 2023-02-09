package commands

import (
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gokeeper/internal/data"
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:       "login",
	Short:     "logins you into the system",
	Args:      cobra.MinimumNArgs(2),
	ValidArgs: []string{"login", "password"},
	Run: func(cmd *cobra.Command, args []string) {
		authorized, err := deps.AuthService.CheckAuthorized(cmd.Context())
		if authorized || errors.Is(err, data.ErrLoggedInAlready) {
			cmd.PrintErrln("You are already logged into the system.")
			return
		}
		if err != nil {
			log.Error().Err(err).Msg("Checking whether authorized user")
			return
		}

		login, password := args[0], args[1]

		if err := deps.UserService.Login(cmd.Context(), login, password); err != nil {
			if errors.Is(err, data.ErrCredentialsDontMatch) {
				cmd.PrintErrln("The credentials don't match any of our records")
				return
			}

			log.Error().Err(err).Msg("Logging user from command")
			return
		}

		cmd.Printf("[%s] Logged in successfully!\n", login)
	},
}
