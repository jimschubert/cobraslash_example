// +build generate

package main

import (
	"fmt"
	"github.com/jimschubert/cobraslash_example/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"os"
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
