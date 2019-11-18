package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"

	catption "github.com/Zenika/catption/lib"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	top, bottom            string
	size, fontSize, margin float64
	out                    = "out.jpg"
	dirs                   = []string{"."}

	version = "chapter8"

	logLevel = logrus.InfoLevel

	cmd = &cobra.Command{
		Use:  "catption",
		Long: "Cat caption generator CLI",
		Args: cobra.ExactArgs(1),
		PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
			logrus.SetLevel(logLevel)

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
			} else {
				logrus.Debugf("Using config file %s", viper.ConfigFileUsed())
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
			logrus.Infof("Found input file %s", path)

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
		Use:  "dir",
		Long: "Adds a directory to the input files directory",
		Args: cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			return addDir(args[0])
		},
	}
)

func init() {
	cmd.Version = version

	cmd.Flags().StringVarP(&top, "top", "t", "", "Top text")
	cmd.Flags().StringVarP(&bottom, "bottom", "b", "", "Bottom text")

	cmd.Flags().Float64VarP(&size, "size", "s", catption.DefaultSize, "Output image size")
	cmd.Flags().Float64Var(&fontSize, "fontSize", catption.DefaultFontSize, "Font in points")
	cmd.Flags().Float64Var(&margin, "margin", catption.DefaultMargin, "Top/bottom text margin")

	cmd.Flags().StringVarP(&out, "out", "o", out, "Output file")

	cmd.Flags().StringSlice("dir", nil, "Input files directory")
	viper.BindPFlag("dirs", cmd.Flags().Lookup("dir"))

	cmd.PersistentFlags().Var((*logLevelValue)(&logLevel), "logLevel", "Log level")

	cmd.AddCommand(dirCmd)
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

type logLevelValue logrus.Level

var _ pflag.Value = new(logLevelValue)

func (l *logLevelValue) Set(value string) error {
	lvl, err := logrus.ParseLevel(value)
	if err != nil {
		return err
	}
	*l = logLevelValue(lvl)
	return nil
}

func (l *logLevelValue) String() string {
	return l.String()
}

func (l *logLevelValue) Type() string {
	return "string"
}
