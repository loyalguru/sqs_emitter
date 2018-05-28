package sqs_emitter

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	s "github.com/aws/aws-sdk-go/service/sqs"
)

type Message struct {
	Body            string
	Group           string
	DeduplicaitonId string
	Attributes      map[string]string
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

	if len(message.Attributes) > 0 {
		var attributes map[string]*s.MessageAttributeValue

		for key, value := range message.Attributes {
			attribute := &s.MessageAttributeValue{}
			attribute.SetDataType("String")
			attribute.SetStringValue(value)

			attributes[key] = attribute
		}

		sendParams.MessageAttributes = attributes
	}

	_, err := e.Client.SendMessage(sendParams)

	if err != nil {
		log.Fatal(err)
	}
}
