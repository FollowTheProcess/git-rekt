package main

import (
	"context"
	"fmt"
	"os"

	"go.followtheprocess.codes/git-rekt/internal/cli"
	"go.followtheprocess.codes/msg"
)

func main() {
	if err := run(); err != nil {
		msg.Error("%s", err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()

	cmd, err := cli.Build()
	if err != nil {
		return fmt.Errorf("could not build git-rekt cli: %w", err)
	}

	return cmd.Execute(ctx)
}
