// +build !js,!wasm

package main

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dirCmd = &cobra.Command{
	Use:  "dir",
	Long: "Adds a directory to the input files directory",
	Args: cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		return addDir(args[0])
	},
}

func init() {
	cmd.AddCommand(dirCmd)
}

func addDir(dir string) error {
	var dirs = viper.GetStringSlice("dirs")

	for _, d := range dirs {
		if d == dir {
			return nil
		}
	}

	viper.Set("dirs", append(dirs, dir))

	if err := viper.WriteConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}

		viper.Set("dirs", []string{dir})

		configDir, err := os.UserConfigDir()
		if err != nil {
			return err
		}

		return viper.WriteConfigAs(filepath.Join(configDir, "catption.json"))
	}

	return nil
}
