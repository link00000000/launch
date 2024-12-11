package launch

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/link00000000/launch/pkg/launch/configurations"
)

type LaunchJSON struct {
	Configurations []json.RawMessage `json:"configurations"`
}

type Launch struct {
	Configurations []configurations.Configuration
}

var ErrConfigurationNotFound = errors.New("configuration not found")

func ReadFromFile(name string) (*Launch, error) {
	b, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	return ReadFromJSON(b)
}

func ReadFromJSON(b []byte) (*Launch, error) {
	var launchJson LaunchJSON

	err := json.Unmarshal(b, &launchJson)
	if err != nil {
		return nil, err
	}

	launch := &Launch{}
	for _, raw := range launchJson.Configurations {
		var base configurations.BaseConfigurationJSON

		err := json.Unmarshal(raw, &base)
		if err != nil {
			return nil, err
		}

		switch base.Type {
		case "go":
			cfg := &configurations.GoConfigurationJSON{}

			err := json.Unmarshal(raw, cfg)
			if err != nil {
				return nil, err
			}

			launch.Configurations = append(launch.Configurations, cfg)
		}
	}

	return launch, nil
}

func (launch *Launch) FindConfiguration(name string) (configurations.Configuration, error) {
	if len(launch.Configurations) == 0 {
		return nil, ErrConfigurationNotFound
	}

	if name == "" {
		return launch.Configurations[0], nil
	}

	for _, cfg := range launch.Configurations {
		if cfg.GetName() == name {
			return cfg, nil
		}
	}

	return nil, ErrConfigurationNotFound
}
