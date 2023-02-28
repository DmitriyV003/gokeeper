package commands

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	rootCmd.AddCommand(deleteCredsCmd)
}

var deleteCredsCmd = &cobra.Command{
	Use:       "delete-creds",
	Short:     "deletes your credentials from the system",
	Args:      cobra.MinimumNArgs(1),
	ValidArgs: []string{"id"},
	Example:   `delete-creds 1`,
	Run: func(cmd *cobra.Command, args []string) {
		authorized, err := deps.AuthService.CheckAuthorized(cmd.Context())
		if err != nil {
			if err := deps.UserService.Delete(cmd.Context()); err != nil {
				log.Error().Err(err).Msg("Deleting user info")
			}
			return
		}
		if !authorized {
			unauthorized(cmd)
			return
		}

		idString := args[0]

		id, err := strconv.Atoi(idString)
		if err != nil {
			cmd.PrintErrf("Unable to convert [%s] to integer", idString)
			return
		}

		if err := deps.LoginSecretService.Delete(cmd.Context(), int64(id)); err != nil {
			log.Error().Err(err).Msg("Deleting creds from command")
			return
		}

		cmd.Println("The credentials have been deleted successfully!")
	},
}
