package saramafx

import (
	"fmt"

	"github.com/IBM/sarama"
)

type SendMessageRequest struct {
	Topic   string
	Message []byte
}

// SendMessage sends a message to the specified kafka topic
func (kc Client) SendMessage(req SendMessageRequest) error {
	msg := &sarama.ProducerMessage{
		Topic: req.Topic,
		Value: sarama.StringEncoder(req.Message),
	}
	partition, offset, err := kc.producer.SendMessage(msg)
	if err != nil {
		return err
	}
	fmt.Printf("Message published to topic(%s)/partition(%d)/offset(%d)\n", req.Topic, partition, offset)
	return nil
}
