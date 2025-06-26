package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"

	"github.com/weanan/weanan-service/cmd/server"
)

func newServerCmd() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "server command handler",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				log.Fatalf("unknown command %q", strings.Join(args, " "))
			}

			_ = cmd.Help()
		},
	}

	serverCmd.AddCommand(server.NewRestServer())
	serverCmd.AddCommand(server.NewSubscriberServer())

	return serverCmd
}

func init() {
	rootCmd.AddCommand(newServerCmd())
}
