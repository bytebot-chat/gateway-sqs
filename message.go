package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/bytebot-chat/gateway-sqs/model"
	"github.com/rs/zerolog/log"
	"github.com/satori/go.uuid"
)

func handleInbound(m *sqs.Message) {
	ctx := context.Background()

	log.Debug().Msg("Handling new message")

	msg := &model.Message{
		Body: *m,
	}
	msg.Metadata.ID = uuid.Must(uuid.NewV4(), *new(error))
	msg.Metadata.Source = *id
	log.Debug().Msg("Marshaling message to JSON...")
	stringMsg, _ := json.Marshal(msg)
	log.Debug().Msg("Publishing...")
	rdb.Publish(ctx, *inbound, stringMsg)
	log.Debug().Msg("Published message to " + *inbound)
	return
}
