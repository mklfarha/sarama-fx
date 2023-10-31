package main

import (
	"fmt"

	saramafx "github.com/mklfarha/sarama-fx"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func main() {
	initConfig()
	fx.New(
		saramafx.Module,
		fx.Invoke(NewKafkaProducer),
	).Run()
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

type kafkaProducer struct {
	client *saramafx.Client
}

func NewKafkaProducer(sc *saramafx.Client) kafkaProducer {
	producer := kafkaProducer{
		client: sc,
	}
	// producing for testing purposes
	producer.Produce()
	return producer
}

func (kp kafkaProducer) Produce() error {
	return kp.client.SendMessage(saramafx.SendMessageRequest{
		Topic:   "test",
		Message: []byte("message from server"),
	})
}
