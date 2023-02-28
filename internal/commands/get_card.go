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
	rootCmd.AddCommand(getCardCmd)
}

var getCardCmd = &cobra.Command{
	Use:       "get-card",
	Short:     "get a specific credentials using ID",
	Args:      cobra.MinimumNArgs(1),
	ValidArgs: []string{"id"},
	Example:   `get-card 1234`,
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

		card, err := deps.CardSecretService.Get(cmd.Context(), int64(id))
		if err != nil {
			log.Error().Err(err).Msg("Getting cardSecret from command")
			return
		}

		displayCardSecret(cmd, card)
	},
}

func displayCardSecret(cmd *cobra.Command, secret *core.CardSecret) {
	if secret == nil {
		cmd.PrintErrln("No login with this ID found")
		return
	}

	line := strings.Repeat("-", 10)
	cmd.Println(fmt.Sprintf("%s %s %s", line, "Card", line))

	cmd.Printf("ID: %d\n", secret.ID)
	cmd.Printf("Cardholder Name: %s\n", secret.CardholderName)
	cmd.Printf("Expire Date: %s\n", secret.ExpireDate)
	cmd.Printf("Valid From: %s\n", secret.ValidFrom)
	cmd.Printf("Secret Code: %s\n", secret.SecretCode)
	cmd.Printf("Number: %s\n", secret.Number)
	cmd.Printf("Additional data: %s\n", secret.AdditionalData)
}
