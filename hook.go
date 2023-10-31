package saramafx

import (
	"context"

	"go.uber.org/fx"
)

// hook to the FX application lifecycle to start consuming
// and close the producer and consumer group on stop
func hook(lifecycle fx.Lifecycle, kc *Client) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// start consuming
			if kc.consumerGroupHandler != nil {
				go Consume(*kc)
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			// close connections
			kc.producer.Close()
			kc.group.Close()
			return nil
		},
	})
}
