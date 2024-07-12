package main

import (
	"time"

	saramafx "github.com/mklfarha/sarama-fx"
	"go.uber.org/config"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			NewConfig,
		),
		saramafx.Module,
		fx.Invoke(NewKafkaProducer),
	).Run()
}

func NewConfig() config.Provider {
	yaml, err := config.NewYAML(config.File("./config.yaml"))
	if err != nil {
		panic("error reading config")
	}
	return yaml
}

type kafkaProducer struct {
	client *saramafx.Client
}

func NewKafkaProducer(sc *saramafx.Client) kafkaProducer {
	producer := kafkaProducer{
		client: sc,
	}
	// producing for testing purposes
	go func() {
		// we need to wait for the prodcuer to be created
		time.Sleep(1 * time.Second)
		producer.Produce()
	}()
	return producer
}

func (kp kafkaProducer) Produce() error {
	return kp.client.SendMessage(saramafx.SendMessageRequest{
		Topic:   "test",
		Message: []byte("message from server"),
	})
}
