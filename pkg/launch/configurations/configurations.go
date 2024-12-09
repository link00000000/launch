package configurations

import (
	"encoding/json"
	"errors"
)

type Configuration interface {
	GetName() string
	Execute(cwd string) (exitCode int, err error)
}

type BaseConfiguration struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

func (cfg *BaseConfiguration) GetName() string {
	return cfg.Name
}

func UnmarshalJSON(b []byte) (Configuration, error) {
	var base BaseConfiguration

	err := json.Unmarshal(b, &base)
	if err != nil {
		return nil, err
	}

	switch base.Type {
	case "go":
		var cfg GoConfiguration

		err := json.Unmarshal(b, &cfg)
		if err != nil {
			return nil, err
		}

		return &cfg, nil

	default:
		return nil, ErrUnsupportedType
	}
}

type ConfigurationExecution struct {
	Configuration *BaseConfiguration
}

var ErrUnsupportedType = errors.New("unsupported type")
var ErrUnsupportedRequest = errors.New("unsupported request")

func NewConfigurationExecution(configuration *BaseConfiguration) *ConfigurationExecution {
	return &ConfigurationExecution{Configuration: configuration}
}
