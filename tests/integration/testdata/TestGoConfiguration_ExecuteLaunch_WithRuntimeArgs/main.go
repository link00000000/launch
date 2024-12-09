package main

import "os"

func main() {
	if len(os.Args) != 4 {
		os.Exit(1)
	}

	if os.Args[1] != "one" {
		os.Exit(1)
	}

	if os.Args[2] != "two" {
		os.Exit(1)
	}

	if os.Args[3] != "three" {
		os.Exit(1)
	}
}
