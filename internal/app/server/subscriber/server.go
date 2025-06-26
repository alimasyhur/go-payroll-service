package subscriber

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/weanan/weanan-service/internal/app/container"
	"github.com/weanan/weanan-service/internal/pkg/logger"
	"google.golang.org/api/option"
)

func StartSubscriberService(container *container.Container) {
	if container == nil {
		panic("container is nil")
	}

	ctx := context.Background()

	projectID := container.Config.GCPPubsub.ProjectID
	sendMessageSubID := container.Config.GCPPubsub.SubscriberSendMessageID

	client, err := pubsub.NewClient(ctx, projectID, option.WithCredentialsJSON([]byte(container.Config.GCPPubsub.ServiceAccount)))
	if err != nil {
		logger.Log.Error(ctx, "Failed to create Pub/Sub client", err)
	}
	defer client.Close()

	sub := client.Subscription(sendMessageSubID)
	SetupSubscriber(ctx, sub, container)
}
