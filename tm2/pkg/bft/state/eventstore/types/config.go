package types

import "github.com/gnolang/gno/tm2/pkg/bft/state/eventstore/null"

// EventStoreParams defines the arbitrary event store config params
type EventStoreParams map[string]any

// Config defines the specific event store configuration
type Config struct {
	EventStoreType string
	Params         EventStoreParams
}

// GetParam fetches the specific config param, if any.
// Returns nil if the param is not present
func (c *Config) GetParam(name string) any {
	if c.Params != nil {
		return c.Params[name]
	}

	return nil
}

// DefaultEventStoreConfig returns the default event store config
func DefaultEventStoreConfig() *Config {
	return &Config{
		EventStoreType: null.EventStoreType,
	}
}
