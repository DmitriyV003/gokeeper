package commands

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCardCmd)
}

var createCardCmd = &cobra.Command{
	Use:       "create-card",
	Short:     "create card secret",
	Args:      cobra.MinimumNArgs(5),
	ValidArgs: []string{"cardholder_name", "type", "expire_date", "valid_from", "number", "secret_code", "additional_data"},
	Run: func(cmd *cobra.Command, args []string) {
		authorized, err := deps.AuthService.CheckAuthorized(cmd.Context())
		if err != nil {
			cmd.PrintErrln("Error to check is authorized")
			return
		}

		if !authorized {
			cmd.PrintErrln("You are not authorized")
			return
		}

		cardholderName, typ, expiredate, validFrom, num, secretCode, addD := args[0], args[1], args[2], args[3], args[4], args[5], args[6]

		if err := deps.CardSecretService.Create(cmd.Context(), cardholderName, typ, expiredate, validFrom, num, secretCode, addD); err != nil {
			if err != nil {
				cmd.PrintErrln("Error to create card secret.")
				return
			}

			return
		}

		cmd.Println("Created new card secret")
	},
}
