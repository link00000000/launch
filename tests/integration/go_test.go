package tests

import (
	"testing"

	"github.com/link00000000/launch/pkg/launch"
	"github.com/link00000000/launch/pkg/launch/configurations"
)

func TestGoConfiguration_ExecuteLaunch(t *testing.T) {
	testEnv := SetupTestEnv(t, "testdata/TestGoConfiguration_ExecuteLaunch")
	defer testEnv.Teardown()

	l, err := launch.ReadFile(".vscode/launch.json")
	if err != nil {
		t.Fatalf("unexpected error while calling launch.ReadFile(\".vscode/launch.json\"): %#v", err)
	}

	cfg, err := l.FindConfiguration("test configuration")
	if err != nil {
		t.Fatalf("unexpected error while calling l.FindConfiguration(\"test configuration\"): %#v", err)
	}

	exitCode, err := cfg.Execute("")
	if err != nil {
		t.Fatalf("unexpected error while calling cfg.Execute(\"test configuration\"): %#v", err)
	}

	if exitCode != 0 {
		t.Fatalf("invalid exit code: expected %d, got %d", 0, exitCode)
	}
}

func TestGoConfiguration_ExecuteLaunch_WithoutProgram(t *testing.T) {
	testEnv := SetupTestEnv(t, "testdata/TestGoConfiguration_ExecuteLaunch_WithoutProgram")
	defer testEnv.Teardown()

	l, err := launch.ReadFile(".vscode/launch.json")
	if err != nil {
		t.Fatalf("unexpected error while calling launch.ReadFile(\".vscode/launch.json\"): %#v", err)
	}

	cfg, err := l.FindConfiguration("test configuration")
	if err != nil {
		t.Fatalf("unexpected error while calling l.FindConfiguration(\"test configuration\"): %#v", err)
	}

	_, err = cfg.Execute("")
	if err == nil {
		t.Fatalf("incorrect error: expected %#v, got nil", configurations.ErrUnsupportedRequest)
	}
}

func TestGoConfiguration_ExecuteLaunch_WithEnv(t *testing.T) {
	testEnv := SetupTestEnv(t, "testdata/TestGoConfiguration_ExecuteLaunch_WithEnv")
	defer testEnv.Teardown()

	l, err := launch.ReadFile(".vscode/launch.json")
	if err != nil {
		t.Fatalf("unexpected error while calling launch.ReadFile(\".vscode/launch.json\"): %#v", err)
	}

	cfg, err := l.FindConfiguration("test configuration")
	if err != nil {
		t.Fatalf("unexpected error while calling l.FindConfiguration(\"test configuration\"): %#v", err)
	}

	exitCode, err := cfg.Execute("")
	if err != nil {
		t.Fatalf("unexpected error while calling cfg.Execute(\"test configuration\"): %#v", err)
	}

	if exitCode != 0 {
		t.Fatalf("invalid exit code: expected %d, got %d", 0, exitCode)
	}
}

func TestGoConfiguration_ExecuteLaunch_WithRuntimeArgs(t *testing.T) {
	testEnv := SetupTestEnv(t, "testdata/TestGoConfiguration_ExecuteLaunch_WithRuntimeArgs")
	defer testEnv.Teardown()

	l, err := launch.ReadFile(".vscode/launch.json")
	if err != nil {
		t.Fatalf("unexpected error while calling launch.ReadFile(\".vscode/launch.json\"): %#v", err)
	}

	cfg, err := l.FindConfiguration("test configuration")
	if err != nil {
		t.Fatalf("unexpected error while calling l.FindConfiguration(\"test configuration\"): %#v", err)
	}

	exitCode, err := cfg.Execute("")
	if err != nil {
		t.Fatalf("unexpected error while calling cfg.Execute(\"test configuration\"): %#v", err)
	}

	if exitCode != 0 {
		t.Fatalf("invalid exit code: expected %d, got %d", 0, exitCode)
	}
}

func TestGoConfiguration_ExecuteUnsupportedRequest(t *testing.T) {
	testEnv := SetupTestEnv(t, "testdata/TestGoConfiguration_ExecuteUnsupportedRequest")
	defer testEnv.Teardown()

	l, err := launch.ReadFile(".vscode/launch.json")
	if err != nil {
		t.Fatalf("unexpected error while calling launch.ReadFile(\".vscode/launch.json\"): %#v", err)
	}

	cfg, err := l.FindConfiguration("test configuration")
	if err != nil {
		t.Fatalf("unexpected error while calling l.FindConfiguration(\"test configuration\"): %#v", err)
	}

	_, err = cfg.Execute("")
	if err != configurations.ErrUnsupportedRequest {
		t.Fatalf("incorrect error: expected %#v, got %#v", configurations.ErrUnsupportedRequest, err)
	}
}
