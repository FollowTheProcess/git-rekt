// Package cli implements git-rekts super helpful command line interface.
package cli

import (
	"context"
	"fmt"
	"os"

	"go.followtheprocess.codes/cli"
	"go.followtheprocess.codes/cli/flag"
	"go.followtheprocess.codes/git-rekt/internal/insults"
)

// These are all set at compile time.
var (
	version   = "dev"
	commit    = ""
	buildDate = ""
)

// Build builds and returns the git-rekt CLI.
func Build() (*cli.Command, error) {
	var (
		hard  bool // If true, insult them really hard
		force bool // Do something *really* mean 👀
	)
	return cli.New(
		"git-rekt",
		cli.Short("A profoundly unhelpful git extension"),
		cli.Version(version),
		cli.Commit(commit),
		cli.BuildDate(buildDate),
		cli.Run(func(ctx context.Context, cmd *cli.Command) error {
			if hard {
				// Really give it to them
				fmt.Println(insults.GetWorse())
			} else {
				// Just a normal insult
				fmt.Println(insults.Get())
			}

			if force {
				if err := os.RemoveAll(".git"); err != nil {
					return fmt.Errorf("oooof you're lucky, your repo was saved by: %w", err)
				}
			}

			return nil
		}),
		cli.Flag(&hard, "hard", flag.NoShortHand, "Do whatever it does... but harder"),
		cli.Flag(&force, "force", flag.NoShortHand, "You won't like this"),
	)
}
