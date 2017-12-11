package sqs_emitter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	s "github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"log"
)

type Config struct {
	QueueName string
	Client    sqsiface.SQSAPI
	Region    string
	QueueUrl  string
}

func (c *Config) defaults() {
	c.Client = s.New(session.New(aws.NewConfig()))

	if c.QueueUrl == "" {
		c.QueueUrl = c.queueUrl(c.QueueName)
	}

	if c.Region == "" {
		c.Region = "eu-west-1"
	}
}

func (c *Config) queueUrl(name string) string {
	fifoQueueName := name

	queueInfo, err := c.Client.GetQueueUrl(&s.GetQueueUrlInput{
		QueueName: &fifoQueueName,
	})

	if err != nil {
		log.Fatalf("Queue name: %v error: %v", fifoQueueName, err)
	}

	return *queueInfo.QueueUrl
}
