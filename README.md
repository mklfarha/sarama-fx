# sarama fx

[![Go Reference](https://pkg.go.dev/badge/github.com/mklfarha/sarama-fx.svg)](https://pkg.go.dev/github.com/mklfarha/sarama-fx)

[FX](https://github.com/uber-go/fx) wrapper for the [sarama](https://github.com/IBM/sarama) go client for [Apache Kafka](https://kafka.apache.org/)

## Getting started

Checkout the **[examples](https://github.com/mklfarha/sarama-fx/tree/main/examples)** directory in the repo.

### Quickstart

Add the module to your FX application

```go
fx.New(		
    saramafx.Module,
).Run()
```

The module uses [viper](https://github.com/spf13/viper) for config management.

It will look for a node called kafka with the following structure:

```yaml
kafka:
  # kafka version running
  version: 3.6.0
  # consumer group id 
  consumer_group_id: test-group
  # list of brokers 
  brokers:
    - localhost:9092
  # list of topics    
  topics:
    - test
```

### Consuming

Provide a [group consumer handler](https://github.com/IBM/sarama/blob/82f0e48b0e2b6dfefc3088fe5e45195f6cc461d8/consumer_group.go#L1072) to consume messages in your application 

```go
fx.New(
    fx.Provide(
        NewKafkaHandler,
    ),
    saramafx.Module,
).Run()

// implement the sarama.ConsumerGroupHandler interface 
// check out the full example under examples/consumer
func NewKafkaHandler() sarama.ConsumerGroupHandler {
    ...
}
```

### Producing 

Inject the saramafx client wherever you need it and use the SendMessage function to produce a message 

```go
// full example under examples/producer
kp.client.SendMessage(saramafx.SendMessageRequest{
		Topic:   "test",
		Message: []byte("message from server"),
	})
```