package config

import "github.com/psiemens/sconfig"

// Config holds the application configuration for an security node.
type Config struct {
	Port                 int `default:"5000"`
	ProcessorQueueBuffer int `default:"1000"`
	ProcessorCacheBuffer int `default:"10000"`
}

// New returns a new Config object.
func New() *Config {
	var conf Config

	err := sconfig.New(&conf).
		FromEnvironment("BAM").
		Parse()

	if err != nil {
		panic(err.Error())
	}

	return &conf
}
