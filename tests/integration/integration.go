package tests

import (
	"embed"
	"fmt"
	"os"
	"testing"
)

//go:embed all:testdata
var testdata embed.FS

type TestEnv struct {
	t           *testing.T
	originalCwd string
}

func SetupTestEnv(t *testing.T, path string) *TestEnv {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal("failed to get working directory", err)
	}

	tmp := t.TempDir()
	t.Log("created temporary directory", tmp)

	err = os.CopyFS(tmp, testdata)
	if err != nil {
		t.Fatalf("failed to copy test environment %s to %s: %#v", path, tmp, err)
	}

	t.Log("unpacked testdata to temporary directory", tmp)

	newCwd := fmt.Sprintf("%s/%s", tmp, path)
	err = os.Chdir(newCwd)
	if err != nil {
		t.Fatalf("failed to change directory to %s: %#v", newCwd, err)
	}

	t.Log("set working directory", newCwd)

	return &TestEnv{t: t, originalCwd: cwd}
}

func (env *TestEnv) Teardown() {
	// Leave the directory so that it can be deleted when the test ends
	err := os.Chdir(env.originalCwd)
	if err != nil {
		env.t.Fatalf("failed to change directory to %s: %#v", env.originalCwd, err)
	}
}
