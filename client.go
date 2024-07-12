package saramafx

import (
	"github.com/IBM/sarama"
	"go.uber.org/config"
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

	Lifecycle      fx.Lifecycle
	ConfigProvider config.Provider
	// ConsumerGroupHandler needs to be provided by the user of the library
	Handler sarama.ConsumerGroupHandler `optional:"true"`
}

// New sarama fx client
func New(params Params) (*Client, error) {
	c := Config{}
	err := params.ConfigProvider.Get("events.transport-config.kafka").Populate(&c)
	if err != nil {
		return nil, err
	}
	kc := Client{
		config:               c,
		consumerGroupHandler: params.Handler,
	}
	return &kc, nil
}
