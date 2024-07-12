package saramafx

// Config for the saramafx client
type Config struct {
	Version         string   `yaml:"version"`
	Brokers         []string `yaml:"brokers"`
	Topics          []string `yaml:"topics"`
	ConsumerGroupID string   `yaml:"consumer_group_id" mapstructure:"consumer_group_id"`
}
