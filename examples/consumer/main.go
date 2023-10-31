package main

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
	saramafx "github.com/mklfarha/sarama-fx"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func main() {
	initConfig()
	fx.New(
		fx.Provide(
			NewKafkaHandler,
		),
		saramafx.Module,
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

type kafkaHandler struct {
}

func NewKafkaHandler() sarama.ConsumerGroupHandler {
	return kafkaHandler{}
}

func (kh kafkaHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}
func (kh kafkaHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (kh kafkaHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Printf("message channel was closed")
				return nil
			}
			log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
