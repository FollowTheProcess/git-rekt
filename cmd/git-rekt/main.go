package main

import (
	"os"

	"github.com/FollowTheProcess/git-rekt/internal/cli"
	"github.com/FollowTheProcess/msg"
)

func main() {
	cmd := cli.Build()
	if err := cmd.Execute(); err != nil {
		msg.Error("%s", err)
		os.Exit(1)
	}
}
