package main

import (
	"os"

	"github.com/mirinsim/test-repo-standard-check/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
