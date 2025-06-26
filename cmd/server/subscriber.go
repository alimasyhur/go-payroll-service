package server

import (
	"github.com/spf13/cobra"

	"github.com/weanan/weanan-service/internal/app/container"
	"github.com/weanan/weanan-service/internal/app/server/subscriber"
)

func NewSubscriberServer() *cobra.Command {
	return &cobra.Command{
		Use:   "subscriber",
		Short: "Run Subscriber Server",
		Long:  "Run Subscriber Server",
		Run: func(cmd *cobra.Command, args []string) {
			container := container.Setup()
			go subscriber.StartSubscriberService(container)
		},
	}
}
