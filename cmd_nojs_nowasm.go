// +build !js,!wasm

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	catption "github.com/nlepage/catption/lib"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	top, bottom            string
	size, fontSize, margin float64
	out                    = "out.jpg"
	dirs                   = []string{"."}
)

func init() {
	cmd.Flags().StringVarP(&top, "top", "t", "", "Top text")
	cmd.Flags().StringVarP(&bottom, "bottom", "b", "", "Bottom text")

	cmd.Flags().Float64VarP(&size, "size", "s", catption.DefaultSize, "Output image size")
	cmd.Flags().Float64Var(&fontSize, "fontSize", catption.DefaultFontSize, "Font in points")
	cmd.Flags().Float64Var(&margin, "margin", catption.DefaultMargin, "Top/bottom text margin")

	cmd.Flags().StringVarP(&out, "out", "o", out, "Output file")

	cmd.Flags().StringSlice("dir", nil, "Input files directory")

	cmd.RunE = func(_ *cobra.Command, args []string) error {
		var name = args[0]

		path, err := resolveName(name)
		if err != nil {
			return err
		}
		logrus.Infof("Found input file %s", path)

		cat, err := catption.LoadJPG(path)
		if err != nil {
			return err
		}

		cat.Top, cat.Bottom = top, bottom
		cat.Size, cat.FontSize, cat.Margin = size, fontSize, margin

		if err := cat.SaveJPG(out); err != nil {
			return err
		}

		return exec.Command(openCmd, out).Run()
	}
}

func resolveName(name string) (string, error) {
	for _, dir := range dirs {
		path := filepath.Join(dir, name)

		stat, err := os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return "", err
		}

		if stat.IsDir() {
			return "", fmt.Errorf("%v is a directory", path)
		}

		return path, nil
	}

	return "", fmt.Errorf("%v not found (dirs=%v)", name, dirs)
}
