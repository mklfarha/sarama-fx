package saramafx

import (
	"errors"

	"github.com/spf13/viper"
)

// Config for the saramafx client
type Config struct {
	Version         string   `yaml:"version"`
	Brokers         []string `yaml:"brokers"`
	Topics          []string `yaml:"topics"`
	ConsumerGroupID string   `yaml:"consumer_group_id" mapstructure:"consumer_group_id"`
}

// parseConfig parses the config using viper,
// expects a tree element named kafak with the structure above
func parseConfig() (Config, error) {
	kafkaViper := viper.Sub("kafka")
	if kafkaViper != nil {
		cfg := Config{}
		kafkaViper.Unmarshal(&cfg)
		return cfg, nil
	}

	return Config{}, errors.New("error parsing saramafx config")
}
