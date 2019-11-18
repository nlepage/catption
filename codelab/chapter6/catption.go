package main

import (
	"fmt"
	"os"
	"path/filepath"

	catption "github.com/Zenika/catption/lib"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	top, bottom            string
	size, fontSize, margin float64
	out                    = "out.jpg"
	dirs                   = []string{"."}

	cmd = &cobra.Command{
		Use:     "catption",
		Long:    "Cat caption generator CLI",
		Args:    cobra.ExactArgs(1),
		Version: "chapter5",
		PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
			viper.SetConfigName("catption")
			viper.AddConfigPath(".")
			if configDir, err := os.UserConfigDir(); err == nil {
				viper.AddConfigPath(configDir)
			}

			viper.AutomaticEnv()
			viper.SetEnvPrefix("catption")

			viper.SetDefault("dirs", dirs)

			if err := viper.ReadInConfig(); err != nil {
				if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
					return err
				}
			}

			dirs = viper.GetStringSlice("dirs")

			return nil
		},
		RunE: func(_ *cobra.Command, args []string) error {
			var name = args[0]

			path, err := resolveName(name)
			if err != nil {
				return err
			}

			cat, err := catption.LoadJPG(path)
			if err != nil {
				return err
			}

			cat.Top, cat.Bottom = top, bottom
			cat.Size, cat.FontSize, cat.Margin = size, fontSize, margin

			return cat.SaveJPG(out)
		},
	}

	dirCmd = &cobra.Command{
		// FIXME Set the following fields: Use, Long, Args and RunE
	}
)

func init() {
	cmd.Flags().StringVarP(&top, "top", "t", "", "Top text")
	cmd.Flags().StringVarP(&bottom, "bottom", "b", "", "Bottom text")

	cmd.Flags().Float64VarP(&size, "size", "s", catption.DefaultSize, "Output image size")
	cmd.Flags().Float64Var(&fontSize, "fontSize", catption.DefaultFontSize, "Font in points")
	cmd.Flags().Float64Var(&margin, "margin", catption.DefaultMargin, "Top/bottom text margin")

	cmd.Flags().StringVarP(&out, "out", "o", out, "Output file")

	cmd.Flags().StringSlice("dir", nil, "Input files directory")
	viper.BindPFlag("dirs", cmd.Flags().Lookup("dir"))

	// FIXME add dirCmd to cmd's subcommands
}

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func addDir(dir string) error {
	viper.Set("dirs", append(viper.GetStringSlice("dirs"), dir))

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
