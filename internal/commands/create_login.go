package commands

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(registerCmd)
}

var createLoginCmd = &cobra.Command{
	Use:       "create-login",
	Short:     "create login secret",
	Args:      cobra.MinimumNArgs(5),
	ValidArgs: []string{"username", "website", "password", "additional_data"},
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

		username, website, password, additionalData := args[0], args[1], args[2], args[3]

		if err := deps.LoginSecretService.Create(cmd.Context(), username, website, password, additionalData); err != nil {
			if err != nil {
				cmd.PrintErrln("Error to create login secret.")
				return
			}

			return
		}

		cmd.Println("Created new login")
	},
}
