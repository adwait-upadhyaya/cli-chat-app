package cmd

import (
	"fmt"
	"log"

	"github.com/adwait-upadhyaya/cli-chat-app/internal/database"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Ya pugyo")
		username := args[0]
		email := args[1]
		password := args[2]

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		if err != nil {
			log.Fatal("Error Hashing Password")
		}

		err = database.RegisterUser(username, email, string(hashedPassword))
		if err != nil {
			log.Fatal("Error registering user", err)
		}

		fmt.Println("User Registered SUccesfully")
	},
}
