package main

import (
	"os"

	"github.com/Zenika/catption/codelab/chapter3/catption"

	"github.com/spf13/cobra"
)

var (
	top, bottom            string
	size, fontSize, margin float64
	out                    = "out.jpg"

	cmd = &cobra.Command{
		Use:     "catption",
		Long:    "Cat caption generator CLI",
		Args:    cobra.ExactArgs(1),
		Version: "chapter4",
		RunE: func(_ *cobra.Command, args []string) error {
			var name = args[0]

			cat, err := catption.LoadJPG(name)
			if err != nil {
				return err
			}

			cat.Top, cat.Bottom = top, bottom
			cat.Size, cat.FontSize, cat.Margin = size, fontSize, margin

			return cat.SaveJPG(out)
		},
	}
)

func init() {
	cmd.Flags().StringVarP(&top, "top", "t", "", "Top text")
	cmd.Flags().StringVarP(&bottom, "bottom", "b", "", "Bottom text")

	cmd.Flags().Float64VarP(&size, "size", "s", catption.DefaultSize, "Output image size")
	cmd.Flags().Float64Var(&fontSize, "fontSize", catption.DefaultFontSize, "Font  (points)")
	cmd.Flags().Float64Var(&margin, "margin", catption.DefaultMargin, "Top/bottom text margin")

	cmd.Flags().StringVarP(&out, "out", "o", out, "Output file")
}

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
