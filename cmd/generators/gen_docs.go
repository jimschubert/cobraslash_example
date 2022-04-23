// +build generate

package main

import (
	"fmt"
	"os"

	"github.com/jimschubert/cobraslash_example/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func main() {
	positional := os.Args[1:]
	docType := "markdown"
	if len(positional) > 0 {
		docType = positional[0]
	}

	fmt.Printf("Generating %s documentation...\n", docType)
	r := cmd.RootCommand()

	switch docType {
	case "yaml":
		cobra.CheckErr(doc.GenYamlTree(&r, "./yaml"))
	case "reStructuredText":
		cobra.CheckErr(doc.GenReSTTree(&r, "./rest"))
	case "manpage":
		header := &doc.GenManHeader{
			Title:   "cobraslash",
			Section: "3",
		}
		cobra.CheckErr(doc.GenManTree(&r, header, "./manpage"))
	case "markdown":
	default:
		cobra.CheckErr(doc.GenMarkdownTree(&r, "./docs"))
	}
}
