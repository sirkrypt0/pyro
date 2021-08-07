package main

import (
	"github.com/sirkrypt0/pyro/cmd/pyro/cmd"
	"os"
)

func main() {
	c := cmd.NewPyroCommand(os.Stdin, os.Stdout, os.Stderr)
	if err := c.Execute(); err != nil {
		os.Exit(1)
	}
}
