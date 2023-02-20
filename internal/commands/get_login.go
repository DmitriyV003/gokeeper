package commands

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gokeeper/internal/core"
	"strconv"
	"strings"
)

func init() {
	rootCmd.AddCommand(getCredsCmd)
}

var getCredsCmd = &cobra.Command{
	Use:       "get-login",
	Short:     "get a specific credentials using ID",
	Args:      cobra.MinimumNArgs(1),
	ValidArgs: []string{"id"},
	Example:   `get-login 1234`,
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

		login, err := deps.LoginSecretService.Get(cmd.Context(), int64(id))
		if err != nil {
			log.Error().Err(err).Msg("Getting loginSecret from command")
			return
		}

		displayLoginSecret(cmd, login)
	},
}

func displayLoginSecret(cmd *cobra.Command, secret *core.LoginSecret) {
	if secret == nil {
		cmd.PrintErrln("No login with this ID found")
		return
	}

	line := strings.Repeat("-", 10)
	cmd.Println(fmt.Sprintf("%s %s %s", line, "Credentials", line))

	cmd.Printf("ID: %d\n", secret.ID)
	cmd.Printf("Website: %s\n", secret.Website)
	cmd.Printf("Username: %s\n", secret.Username)
	cmd.Printf("Password: %s\n", secret.Password)
	cmd.Printf("Additional data: %s\n", secret.AdditionalData)
}
