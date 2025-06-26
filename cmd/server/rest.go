package server

import (
	"github.com/spf13/cobra"

	"github.com/weanan/weanan-service/internal/app/container"
	"github.com/weanan/weanan-service/internal/app/server/rest"
	"github.com/weanan/weanan-service/internal/app/server/subscriber"
)

func NewRestServer() *cobra.Command {
	return &cobra.Command{
		Use:   "rest",
		Short: "Run Rest Http Server",
		Long:  "Run Rest Http Server",
		Run: func(cmd *cobra.Command, args []string) {
			container := container.Setup()
			// subscriber.StartSubscriberService(container)
			go subscriber.StartSubscriberService(container)
			rest.StartRestHttpService(container)

		},
	}
}
