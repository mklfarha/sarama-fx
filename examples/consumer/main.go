package main

import (
	"log"

	"github.com/IBM/sarama"
	saramafx "github.com/mklfarha/sarama-fx"
	"go.uber.org/config"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			NewKafkaHandler,
			NewConfig,
		),
		saramafx.Module,
	).Run()
}

func NewConfig() config.Provider {
	yaml, err := config.NewYAML(config.File("./config.yaml"))
	if err != nil {
		panic("error reading config")
	}
	return yaml
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
