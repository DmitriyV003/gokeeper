package commands

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(logoutCmd)
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "logs you out of the system",
	Run: func(cmd *cobra.Command, args []string) {
		authorized, _ := deps.AuthService.CheckAuthorized(cmd.Context())
		if !authorized {
			unauthorized(cmd)
			return
		}

		if err := deps.UserService.Delete(cmd.Context()); err != nil {
			log.Error().Err(err).Msg("Logging user from command")
			return
		}

		cmd.Println("Now you logged out.")
	},
}
