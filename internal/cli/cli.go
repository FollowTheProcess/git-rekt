// Package cli implements git-rekts super helpful command line interface.
package cli

import (
	"fmt"
	"os"

	"github.com/FollowTheProcess/git-rekt/internal/insults"
	"github.com/spf13/cobra"
)

// These are all set at compile time.
var (
	version   = "dev"
	commit    = ""
	buildDate = ""
	builtBy   = ""
)

// Build builds and returns the git-rekt CLI.
func Build() *cobra.Command {
	var (
		hard  bool // If true, insult them really hard
		force bool // Do something *really* mean 👀
	)
	cmd := &cobra.Command{
		Use:           "git-rekt",
		Version:       version,
		Args:          cobra.NoArgs,
		SilenceUsage:  true,
		SilenceErrors: true,
		Short:         "A profoundly unhelpful git extension",
		RunE: func(cmd *cobra.Command, args []string) error {
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
		},
	}
	// Set our custom version and usage templates
	cmd.SetUsageTemplate(usageTemplate)
	cmd.SetVersionTemplate(versionTemplate)

	// Set the flags
	cmd.Flags().BoolVarP(&hard, "hard", "", false, "Do whatever it does... but harder")
	cmd.Flags().BoolVarP(&force, "force", "f", false, "You won't like this")

	return cmd
}
