// +build generate

package main

import (
	"github.com/jimschubert/cobraslash_example/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func main() {
	println("Generating documentation...")
	r := cmd.RootCommand()
	cobra.CheckErr(doc.GenMarkdownTree(&r, "./docs"))
}
