package cmd

import (
	"log"

	"github.com/adwait-upadhyaya/cli-chat-app/internal/database"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login a user",
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		password := args[1]

		err := database.LoginUser(username, password)
		if err != nil {
			log.Fatal("Error occured during login")
		}
	},
}
