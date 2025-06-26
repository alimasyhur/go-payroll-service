package subscriber

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/weanan/weanan-service/internal/app/container"
	"github.com/weanan/weanan-service/internal/app/handler/subscriber/organization"
	"github.com/weanan/weanan-service/internal/pkg/logger"
)

func SetupSubscriber(ctx context.Context, sub *pubsub.Subscription, container *container.Container) {
	organizationHandler := organization.NewHandler().SetOrganizationUsecase(container.OrganizationUsecase).Validate()
	err := sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		logger.Log.Info(ctx, "Received message", string(msg.Data))
		organizationHandler.Create(ctx, *msg)

		msg.Ack()
	})

	if err != nil {
		logger.Log.Error(ctx, "Failed to receive messages", err)
	}
}
