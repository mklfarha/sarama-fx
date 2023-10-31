package saramafx

import (
	"time"

	"github.com/IBM/sarama"
	"go.uber.org/fx"
)

// Client for sarama fx
type Client struct {
	config               Config
	producer             sarama.SyncProducer
	group                sarama.ConsumerGroup
	consumerGroupHandler sarama.ConsumerGroupHandler
}

// Params to create the client
type Params struct {
	fx.In

	Lifecycle fx.Lifecycle
	Config    Config
	// ConsumerGroupHandler needs to be provided by the user of the library
	Handler sarama.ConsumerGroupHandler `optional:"true"`
}

// New sarama fx client
func New(params Params) (*Client, error) {
	// TODO: make the timeout configurable
	producer, err := newProducerWithTimeout(params.Config, 2*time.Second)
	if err != nil {
		return nil, err
	}

	// TODO: make the timeout configurable
	group, err := newConsumerGroupWithTimeout(params.Config, 2*time.Second)
	if err != nil {
		return nil, err
	}

	kc := Client{
		config:               params.Config,
		producer:             producer,
		group:                group,
		consumerGroupHandler: params.Handler,
	}
	return &kc, nil
}
