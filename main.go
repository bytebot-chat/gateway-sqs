package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/bytebot-chat/sdk-go"

	// TODO: Wrap these in the sdk-go library
	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
)

// Variables used for command line parameters
var (
	Token string
	ctx   context.Context
	rdb   *redis.Client

	redisAddr = flag.String("redis", "localhost:6379", "Address and port of redis host")
	queueName = flag.String("q", "", "AWS SQS queue name to subscribe to")
	timeout   = flag.Int64("t", 5, "How long, in seconds, that the message is hidden from others")
	region    = flag.String("region", "us-east-1", "AWS Region")
	id        = flag.String("id", "sns", "ID to use when publishing messages")
	inbound   = flag.String("inbound", "sns-inbound", "Pubsub queue to publish inbound messages to")
	outbound  = flag.String("outbound", fmt.Sprintf(*id+"-outbound"), "Pubsub to subscribe to for sending outbound messages.")
)

func init() {
	var err error
	flag.Parse()

	rdb, err = bytebot.Connect(*redisAddr)
	if err != nil {
		log.Fatal().Err(err)
		os.Exit(1)
	}
}

func main() {
	log.Info().
		Str("Redis address", *redisAddr).
		Msg("Bytebot SQS Gateway starting up!")

	svc := NewSQSSession(*region)

	err := subscribeSQS(svc, queueName)
	if err != nil {
		fmt.Println(err)
		log.Error().Err(err)
		os.Exit(1)
	}
}
