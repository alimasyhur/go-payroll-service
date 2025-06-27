package cmd

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/alimasyhur/go-payroll-service/internal/app/container"
	"github.com/alimasyhur/go-payroll-service/internal/app/handler/rest/auth"
	"github.com/alimasyhur/go-payroll-service/internal/pkg/logger"
	"github.com/spf13/cobra"
)

var generateUserCmd = &cobra.Command{
	Use:   "user-seed",
	Short: "Generate User Seeder to DB",
	Run: func(cmd *cobra.Command, args []string) {
		ctr := container.Setup()
		userHandler := auth.NewHandler().
			SetAuthUsecase(ctr.UserUsecase).Validate()

		ctx := context.Background()
		// atau
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := userHandler.GenerateUserSeed(ctx)
		if err != nil {
			logger.Log.Info(cmd.Context(), err.Error())
			return
		}

		logger.Log.Info(cmd.Context(), "Successful start backup daily feeder feed")
	},
}

func newSeederCmd() *cobra.Command {
	seederCmd := &cobra.Command{
		Use:   "seeder",
		Short: "Generate User Seeder to DB",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				log.Fatalf("unknown command %q", strings.Join(args, " "))
			}

			_ = cmd.Help()
		},
	}

	seederCmd.AddCommand(generateUserCmd)

	return seederCmd
}

func init() {
	rootCmd.AddCommand(newSeederCmd())
}
