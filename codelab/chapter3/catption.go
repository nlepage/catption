package main

import (
	"os"

	"github.com/Zenika/catption/codelab/chapter3/catption"

	"github.com/spf13/cobra"
)

var (
	top, bottom            string
	size, fontSize, margin float64

	cmd = &cobra.Command{
		Use:     "catption",
		Long:    "Cat caption generator CLI",
		Args:    cobra.ExactArgs(1),
		Version: "chapter3",
		RunE: func(_ *cobra.Command, args []string) error {
			var name = args[0]

			cat, err := catption.LoadJPG(name)
			if err != nil {
				return err
			}

			cat.Top, cat.Bottom = top, bottom
			cat.Size, cat.FontSize, cat.Margin = size, fontSize, margin

			return cat.SaveJPG("out.jpg")
		},
	}
)

func init() {
	// FIXME configure text flags for top and bottom
	// FIXME configure float flags for size, fontSize and margin
}

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
