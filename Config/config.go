/* Package config is used to store the configurations for the consensus algorithm
* */
package config

import (
	"math/rand"
	"time"

	shared "consensus.com/shared"
)

type ConfigEnv struct {
	MinTimeout time.Duration
	MaxTimeout time.Duration
	Peers      []shared.NodeEntry
	Majority   int
}

func (c *ConfigEnv) RandomizedTimeout() time.Duration {
	if c.MaxTimeout <= c.MinTimeout {
		return c.MinTimeout
	}

	window := c.MaxTimeout - c.MinTimeout + 1

	randomOffset := rand.Int63n(int64(window))

	return c.MinTimeout + time.Duration(randomOffset)
}
