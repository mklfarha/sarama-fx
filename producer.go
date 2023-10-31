package saramafx

import (
	"errors"
	"time"

	"github.com/IBM/sarama"
)

// newProducerWithTimeout new sync sarama producer
func newProducerWithTimeout(cfg Config, timeout time.Duration) (sarama.SyncProducer, error) {
	result := make(chan producerResult, 1)
	go func() {
		result <- newProducer(cfg)
	}()
	select {
	case <-time.After(timeout):
		return nil, errors.New("new kafka producer timed out")
	case result := <-result:
		return result.producer, result.err
	}
}

type producerResult struct {
	producer sarama.SyncProducer
	err      error
}

func newProducer(cfg Config) producerResult {
	kfversion, err := sarama.ParseKafkaVersion(cfg.Version)
	if err != nil {
		return producerResult{
			producer: nil,
			err:      err,
		}
	}
	config := sarama.NewConfig()
	config.Version = kfversion
	config.Producer.Idempotent = true
	config.Producer.Return.Errors = true
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	config.Producer.Transaction.Retry.Backoff = 10
	config.Producer.Retry.Max = 5
	config.Net.MaxOpenRequests = 1

	producer, err := sarama.NewSyncProducer(cfg.Brokers, config)
	if err != nil {
		return producerResult{
			producer: nil,
			err:      err,
		}
	}

	return producerResult{
		producer: producer,
		err:      err,
	}
}
