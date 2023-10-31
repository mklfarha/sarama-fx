package saramafx

import (
	"context"
	"log"
)

// Consume background thread for the consumer to run
func Consume(kc Client) {
	ctx := context.Background()
	for {
		topics := kc.config.Topics
		err := kc.group.Consume(ctx, topics, kc.consumerGroupHandler)
		if err != nil {
			log.Println(err)
		}
	}
}
