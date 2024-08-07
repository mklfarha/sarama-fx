package saramafx

import (
	"go.uber.org/fx"
)

// Module srarma fx module to be provided
var Module = fx.Module("saramafx",
	fx.Provide(
		// create the sarama fx client
		New,
	),

	// hooks to the FX application lifecycle to start consuming
	fx.Invoke(hook),
)
