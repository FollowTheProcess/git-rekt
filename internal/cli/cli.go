// Package cli implements git-rekts super helpful command line interface.
package cli

import (
	"fmt"

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
	cmd := &cobra.Command{
		Use:           "git rekt [OPTIONS]",
		Version:       version,
		Args:          cobra.NoArgs,
		SilenceUsage:  true,
		SilenceErrors: true,
		Short:         "A profoundly unhelpful git extension",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Do something!")
			return nil
		},
	}
	// Set our custom version and usage templates
	cmd.SetUsageTemplate(usageTemplate)
	cmd.SetVersionTemplate(versionTemplate)

	return cmd
}
