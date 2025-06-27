package server

import (
	"github.com/spf13/cobra"

	"github.com/alimasyhur/go-payroll-service/internal/app/container"
	"github.com/alimasyhur/go-payroll-service/internal/app/server/rest"
)

func NewRestServer() *cobra.Command {
	return &cobra.Command{
		Use:   "rest",
		Short: "Run Rest Http Server",
		Long:  "Run Rest Http Server",
		Run: func(cmd *cobra.Command, args []string) {
			container := container.Setup()
			rest.StartRestHttpService(container)

		},
	}
}
