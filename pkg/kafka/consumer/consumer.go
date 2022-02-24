package consumer

import (
	"context"
	"fmt"
	"log"

	"github.com/pcpratheesh/golang-kafka-realtime-data-pipline/config"
	"github.com/segmentio/kafka-go"
)

func Consume(ctx context.Context) {
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: config.GetConfig().KafkaInstance.Brokers,
		Topic:   config.GetConfig().KafkaInstance.Topic,
		GroupID: config.GetConfig().KafkaInstance.GroupID,
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			log.Fatalf("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))
	}
}
