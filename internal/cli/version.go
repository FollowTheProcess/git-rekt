package cli

import (
	"fmt"

	"github.com/fatih/color"
)

// Style for version "headers".
var versionStyle = color.New(color.Bold)

var versionTemplate = fmt.Sprintf(
	`{{printf "%s %s\n%s %s\n%s %s\n%s %s\n"}}`,
	versionStyle.Sprint("Version:"),
	version,
	versionStyle.Sprint("Commit:"),
	commit,
	versionStyle.Sprint("Build Date:"),
	buildDate,
	versionStyle.Sprint("Built By:"),
	builtBy,
)
