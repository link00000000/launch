package launch

import (
	"testing"

	"github.com/link00000000/launch/pkg/launch/configurations"
)

func TestReadFromJSON(t *testing.T) {
	b := []byte(`
    {
      "version": 2.0,
      "configurations": []
    }
  `)

	launch, err := ReadFromJSON(b)

	if err != nil {
		t.Fatal("launch.UnmarshalJSON: ", err)
	}

	if len(launch.Configurations) != 0 {
		t.Fatalf("assert: len(l.Configurations) == 0, got %d", len(launch.Configurations))
	}
}

func TestFindConfiguration(t *testing.T) {
	launch := Launch{
		Configurations: []configurations.Configuration{
			configurations.NewDummyConfiguration("test configuration 1"),
			configurations.NewDummyConfiguration("test configuration 2"),
			configurations.NewDummyConfiguration("test configuration 3"),
		},
	}

	cfg, err := launch.FindConfiguration("test configuration 2")

	if err != nil {
		t.Fatalf("failed to find configuration: %#v", err)
	}

	if cfg.GetName() != "test configuration 2" {
		t.Fatalf("incorrect configuration name: expected %s, got %s", "test configuration 2", cfg.GetName())
	}
}

func TestFindConfiguration_DefaultToFirstConfiguration(t *testing.T) {
	launch := Launch{
		Configurations: []configurations.Configuration{
			configurations.NewDummyConfiguration("test configuration 1"),
			configurations.NewDummyConfiguration("test configuration 2"),
			configurations.NewDummyConfiguration("test configuration 3"),
		},
	}

	cfg, err := launch.FindConfiguration("")

	if err != nil {
		t.Fatalf("failed to find configuration: %#v", err)
	}

	if cfg.GetName() != "test configuration 1" {
		t.Fatalf("incorrect configuration name: expected %s, got %s", "test configuration 1", cfg.GetName())
	}
}

func TestFindConfiguration_ConfigurationDoesNotExist(t *testing.T) {
	launch := Launch{
		Configurations: []configurations.Configuration{
			configurations.NewDummyConfiguration("test configuration 1"),
			configurations.NewDummyConfiguration("test configuration 2"),
			configurations.NewDummyConfiguration("test configuration 3"),
		},
	}

	_, err := launch.FindConfiguration("test configuration 4")

	if err == nil {
		t.Fatalf("did not get expected error ErrConfigurationNotFound")
	}
}

func TestFindConfiguration_NoConfigurationsSpecified(t *testing.T) {
	launch := Launch{
		Configurations: []configurations.Configuration{},
	}

	_, err := launch.FindConfiguration("")

	if err == nil {
		t.Fatalf("did not get expected error ErrConfigurationNotFound")
	}
}
