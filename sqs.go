package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/rs/zerolog/log"
)

func NewSQSSession(region string) *sqs.SQS {
	config := aws.Config{
		Region: aws.String(region),
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config:            config,
	}))

	return sqs.New(sess)
}

func subscribeSQS(svc *sqs.SQS, name *string) error {

	urlResult, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: name,
	})

	if err != nil {
		return err
	}

	queueURL := urlResult.QueueUrl

	for {
		msgResult, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
			AttributeNames: []*string{
				aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
			},
			MessageAttributeNames: []*string{
				aws.String(sqs.QueueAttributeNameAll),
			},
			QueueUrl:            queueURL,
			MaxNumberOfMessages: aws.Int64(1),
			VisibilityTimeout:   timeout,
		})
		if err != nil {
			log.Error().Err(err)
		}

		if len(msgResult.Messages) > 0 {
			fmt.Println("Message Handle: " + *msgResult.Messages[0].ReceiptHandle)
			deleteMessage(svc, queueURL, msgResult.Messages[0].ReceiptHandle)
		}
	}

}

func deleteMessage(svc *sqs.SQS, queueURL, messageHandle *string) error {
	_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      queueURL,
		ReceiptHandle: messageHandle,
	})
	return err
}
