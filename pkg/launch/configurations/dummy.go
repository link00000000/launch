package configurations

type DummyConfiguration struct {
	*BaseConfigurationJSON
}

func NewDummyConfiguration(name string) *DummyConfiguration {
	return &DummyConfiguration{&BaseConfigurationJSON{Name: name}}
}

// Implements [Configuration]
func (DummyConfiguration) Execute(cwd string) (exitCode int, err error) {
	return 0, nil
}
