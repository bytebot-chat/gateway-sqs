package model

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/bytebot-chat/sdk-go"
)

type Message struct {
	Body     sqs.Message
	Metadata bytebot.Metadata // Bytebot message metadata
}

func (m *Message) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

func (m *Message) Unmarshal(b []byte) error {
	if err := json.Unmarshal(b, m); err != nil {
		return err
	}
	return nil
}
