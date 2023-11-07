package saramafx

import (
	"context"
	"log"
	"time"

	"go.uber.org/fx"
)

// hook to the FX application lifecycle to start consuming
// and close the producer and consumer group on stop
func hook(lifecycle fx.Lifecycle, kc *Client) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go initProducer(kc)
			go initGroup(kc)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			// close connections
			if kc.producer != nil {
				kc.producer.Close()
			}
			if kc.group != nil {
				kc.group.Close()
			}
			return nil
		},
	})
}

func initProducer(kc *Client) {
	for kc.producer == nil {
		// TODO: make the timeout configurable
		producer, err := newProducerWithTimeout(kc.config, 2*time.Second)
		if err == nil && producer != nil {
			kc.producer = producer
		} else {
			log.Default().Printf("error creating sarama fx producer: %v", err)
		}
	}
}

func initGroup(kc *Client) {
	for kc.group == nil {
		// TODO: make the timeout configurable
		group, err := newConsumerGroupWithTimeout(kc.config, 2*time.Second)
		if err == nil && group != nil {
			kc.group = group
		} else {
			log.Default().Printf("error creating sarama fx consumer group: %v", err)
		}
	}

	// start consuming
	if kc.consumerGroupHandler != nil {
		go Consume(*kc)
	}
}
