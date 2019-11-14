package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	// FIXME fill Use and Long fields
	RunE: func(_ *cobra.Command, args []string) error {
		return nil
	},
}

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func sayHello(args []string) error {
	if _, err := fmt.Printf("Hello %s!\n", strings.Join(args, " ")); err != nil {
		return err
	}
	return nil
}
