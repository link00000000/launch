package configurations

import "testing"

func TestGoConfigurationUnmarshalJSON(t *testing.T) {
	b := []byte(`
    {
      "name": "test configuration",
      "type": "go",
      "request": "launch",
      "program": "main.go",
      "args": ["--help"],
      "env": {
        "TEST_ONE": "1",
        "TEST_TWO": "custom"
      }
    }
  `)

	cfg, err := UnmarshalJSON(b)

	if err != nil {
		t.Fatal("configurations.UnmarshalJSON: ", err)
	}

	goCfg, ok := cfg.(*GoConfiguration)
	if !ok {
		t.Fatal("configuration not parsed as type GoConfiguration")
	}

	if goCfg.Name != "test configuration" {
		t.Fatalf("incorrect value for property goCfg.Name: expected %s, got %s", "test configuration", goCfg.Name)
	}

	if goCfg.Type != "go" {
		t.Fatalf("incorrect value for property goCfg.Type: expected %s, got %s", "go", goCfg.Type)
	}

	if goCfg.Request != "launch" {
		t.Fatalf("incorrect value for property goCfg.Request: expected %s, got %s", "launch", goCfg.Request)
	}

	if goCfg.Program != "main.go" {
		t.Fatalf("incorrect value for property goCfg.Program: expected %s, got %s", "main.go", goCfg.Program)
	}

	if len(goCfg.Args) != 1 && goCfg.Args[0] != "--help" {
		t.Fatalf("incorrect value for property goCfg.Args: expected %#v, got %#v", []string{"--help"}, goCfg.Args)
	}

	if v, ok := goCfg.Env["TEST_ONE"]; !ok {
		t.Fatalf("no value set for goCfg.Env[\"TEST_ONE\"]")
	} else if v != "1" {
		t.Fatalf("incorrect value for property goCfg.Env[\"TEST_ONE\"]: expected %s, got %s", "1", v)
	}

	if v, ok := goCfg.Env["TEST_TWO"]; !ok {
		t.Fatalf("no value set for goCfg.Env[\"TEST_TWO\"]")
	} else if v != "custom" {
		t.Fatalf("incorrect value for property goCfg.Env[\"TEST_TWO\"]: expected %s, got %s", "custom", v)
	}
}
