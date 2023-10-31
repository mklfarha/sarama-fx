package saramafx

import (
	"errors"
	"time"

	"github.com/IBM/sarama"
)

// newConsumerGroupWithTimeout create a new sarama consumer group from client
func newConsumerGroupWithTimeout(cfg Config, timeout time.Duration) (sarama.ConsumerGroup, error) {
	client, err := newSaramaClientWithTimeout(cfg, timeout)
	if err != nil {
		return nil, err
	}

	group, err := sarama.NewConsumerGroupFromClient(cfg.ConsumerGroupID, client)
	if err != nil {
		return nil, err
	}
	return group, nil
}

// newSaramaClientWithTimeout create a new sarama client
func newSaramaClientWithTimeout(cfg Config, timeout time.Duration) (sarama.Client, error) {
	result := make(chan clientResult, 1)
	go func() {
		result <- newSaramaClient(cfg)
	}()
	select {
	case <-time.After(timeout):
		return nil, errors.New("new kafka client timed out")
	case result := <-result:
		return result.client, result.err
	}
}

type clientResult struct {
	client sarama.Client
	err    error
}

func newSaramaClient(cfg Config) clientResult {
	kfversion, err := sarama.ParseKafkaVersion(cfg.Version)
	if err != nil {
		return clientResult{
			client: nil,
			err:    err,
		}
	}

	config := sarama.NewConfig()
	config.Version = kfversion
	config.Consumer.Return.Errors = true
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}

	client, err := sarama.NewClient(cfg.Brokers, config)
	if err != nil {
		return clientResult{
			client: nil,
			err:    err,
		}
	}

	return clientResult{
		client: client,
		err:    nil,
	}
}
