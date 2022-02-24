package main

import (
	"context"
	"log"

	"github.com/pcpratheesh/golang-kafka-realtime-data-pipline/config"
	"github.com/pcpratheesh/golang-kafka-realtime-data-pipline/pkg/kafka/consumer"
)

func main() {
	// load configurations
	_, err := config.LoadConfiguration()
	if err != nil {
		log.Fatalf("unable to load configurations %v", err)
	}

	// create a new context
	ctx := context.Background()
	consumer.Consume(ctx)
}
