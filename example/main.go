package main

import (
	"github.com/javifr/sqs_emitter"
)

func main() {
	emi := sqs_emitter.New(
		sqs_emitter.Config{
			QueueName: "test",
		},
	)

	emi.Put(&sqs_emitter.Message{
		Body:  "test",
		Group: "post",
	})
}
