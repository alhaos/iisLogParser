package configuration

import (
	"encoding/json"
	"fmt"
	"os"
)

// NewConfig config constructor
func NewConfig(configFilename string) (*Config, error) {

	c := &Config{}

	data, err := os.ReadFile(configFilename)
	if err != nil {
		return nil, fmt.Errorf("unable open configuration file: %w", err)
	}

	err = json.Unmarshal(data, c)
	if err != nil {
		return nil, fmt.Errorf("unable unmarshal configuration file: %w", err)
	}

	return c, nil
}
