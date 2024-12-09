package main

import (
	"flag"
	"log"
	"os"

	"github.com/link00000000/launch/pkg/launch"
)

var (
	launchJsonPath string
	configName     string
)

func init() {
	flag.StringVar(&launchJsonPath, "launch-json", ".vscode/launch.json", "Path to launch.json")
	flag.StringVar(&configName, "configuration", "", "Configuration to run. Leaving this blank will default to the first configuration found")

	flag.Parse()
}

func main() {
	launch, err := launch.ReadFile(".vscode/launch.json")
	if err != nil {
		log.Fatalf("failed to read %s: %#v\n", launchJsonPath, err)
	}

	cfg, err := launch.FindConfiguration(configName)
	if err != nil {
		log.Fatalf("failed to find configuration %s: %#v\n", configName, err)
	}

	exitCode, err := cfg.Execute("")
	if err != nil {
		log.Fatalf("failed to execute configuration %s: %#v", configName, err)
	}

	os.Exit(exitCode)
}
