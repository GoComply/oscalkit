package main

import (
	"fmt"
	"os"

	"github.com/gocomply/oscalkit/cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Troubles. %v", err)
	}
}
