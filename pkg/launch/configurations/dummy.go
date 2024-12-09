package configurations

type DummyConfiguration struct {
	*BaseConfiguration
}

func NewDummyConfiguration(name string) *DummyConfiguration {
	return &DummyConfiguration{&BaseConfiguration{Name: name}}
}

// Implements [Configuration]
func (DummyConfiguration) Execute(cwd string) (exitCode int, err error) {
	return 0, nil
}
