package cmd

import (
	"github.com/adwait-upadhyaya/cli-chat-app/internal/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start socket server",
	Run: func(cmd *cobra.Command, args []string) {
		server.InitServer()
	},
}
