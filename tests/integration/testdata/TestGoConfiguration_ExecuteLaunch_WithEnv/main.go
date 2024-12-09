package main

import "os"

func main() {
	if v, ok := os.LookupEnv("ENV_ONE"); !ok || v != "one" {
		os.Exit(1)
	}

	if v, ok := os.LookupEnv("ENV_TWO"); !ok || v != "two" {
		os.Exit(1)
	}

	if v, ok := os.LookupEnv("ENV_THREE"); !ok || v != "three" {
		os.Exit(1)
	}
}
