package sqs_emitter

import (
	"github.com/aws/aws-sdk-go/aws"
	s "github.com/aws/aws-sdk-go/service/sqs"
	"log"
)

type Message struct {
	Body            string
	Group           string
	DeduplicaitonId string
}

type Emitter struct {
	Config
}

func New(config Config) *Emitter {
	config.defaults()
	emitter := &Emitter{
		Config: config,
	}

	return emitter
}

func (e *Emitter) Put(message *Message) {
	sendParams := &s.SendMessageInput{
		MessageBody: aws.String(message.Body),
		QueueUrl:    aws.String(e.QueueUrl),
	}

	if message.DeduplicaitonId != "" {
		sendParams.MessageDeduplicationId = aws.String(message.DeduplicaitonId)
	}

	if message.Group != "" {
		sendParams.MessageGroupId = aws.String(message.Group)
	}

	_, err := e.Client.SendMessage(sendParams)

	if err != nil {
		log.Fatal(err)
	}
}
