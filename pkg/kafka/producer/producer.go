package producer

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/pcpratheesh/golang-kafka-realtime-data-pipline/config"
	"github.com/segmentio/kafka-go"
)

func Produce(ctx context.Context) {
	// initialize a counter
	i := 0

	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: config.GetConfig().KafkaInstance.Brokers,
		Topic:   config.GetConfig().KafkaInstance.Topic,
	})

	for {
		// each kafka message has a key and value. The key is used
		// to decide which partition (and consequently, which broker)
		// the message gets published on
		err := w.WriteMessages(ctx, kafka.Message{
			Key: []byte(strconv.Itoa(i)),
			// create an arbitrary message payload for the value
			Value: []byte("this is message" + strconv.Itoa(i)),
		})
		if err != nil {
			log.Fatalf("could not write message " + err.Error())
		}

		// log a confirmation once the message is written
		fmt.Println("writes:", i)
		i++
		// sleep for a second
		time.Sleep(time.Second)
	}
}
