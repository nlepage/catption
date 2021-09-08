package main

import (
	"fmt"
	"mime"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	catption "github.com/nlepage/catption/lib"
)

var (
	top, bottom            string
	size, fontSize, margin float64
	out                    = "out.jpg"
	dirs                   []string
	open                   = true

	version = "master"

	logLevel = logrus.InfoLevel

	cmd = &cobra.Command{
		Use:   "catption [flags] <input file>",
		Short: "Cat caption generator CLI",
		Args:  cobra.ExactArgs(1),
		PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
			logrus.SetLevel(logLevel)

			viper.SetConfigName("catption")
			viper.AddConfigPath(".")
			if configDir, err := os.UserConfigDir(); err == nil {
				viper.AddConfigPath(configDir)
			}

			viper.AutomaticEnv()
			viper.SetEnvPrefix("catption")

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

			logrus.Infof("Writing to %s", out)

			if err := cat.SaveJPG(out); err != nil {
				return err
			}

			if open {
				return exec.Command(openCmd, out).Start()
			}

			return nil
		},
	}

	dirCmd = &cobra.Command{
		Use:   "dir",
		Short: "Manages input files directories",
	}

	dirAddCmd = &cobra.Command{
		Use:   "add <directory>",
		Short: "Adds an input files directory",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			return addDir(args[0])
		},
	}

	dirRemoveCmd = &cobra.Command{
		Use:   "remove <directory>",
		Short: "Removes an input files directory",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			return removeDir(args[0])
		},
	}

	dirListCmd = &cobra.Command{
		Use:   "list",
		Short: "Lists input files directories",
		Args:  cobra.NoArgs,
		Run: func(_ *cobra.Command, args []string) {
			listDirs()
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
	cmd.Flags().BoolVar(&open, "open", open, "Open file with system viewer")

	viper.BindPFlag("dirs", cmd.Flags().Lookup("dir"))

	cmd.PersistentFlags().Var((*logLevelValue)(&logLevel), "logLevel", "Log level")

	cmd.AddCommand(dirCmd)
	dirCmd.AddCommand(dirAddCmd)
	dirCmd.AddCommand(dirRemoveCmd)
	dirCmd.AddCommand(dirListCmd)
}

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func addDir(dir string) error {
	for _, d := range dirs {
		if d == dir {
			return nil
		}
	}

	if len(dirs) == 1 && dirs[0] == "." {
		dirs = []string{dir}
	} else {
		dirs = append(dirs, dir)
	}

	viper.Set("dirs", dirs)

	if err := viper.WriteConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}

		configDir, err := os.UserConfigDir()
		if err != nil {
			return err
		}

		configFile := filepath.Join(configDir, "catption.json")

		logrus.Infof("Creating config file %s", configFile)

		return viper.WriteConfigAs(configFile)
	}

	return nil
}

func removeDir(dir string) error {
	var index = -1
	for i, d := range dirs {
		if d == dir {
			index = i
			break
		}
	}

	if index == -1 {
		return nil
	}

	dirs = append(dirs[:index], dirs[index+1:]...)

	viper.Set("dirs", dirs)

	return viper.WriteConfig()
}

func listDirs() {
	for _, d := range dirs {
		fmt.Println(d)
	}
}

func resolveName(name string) (string, error) {
	names := []string{name}

	if ext := filepath.Ext(name); ext == "" || mime.TypeByExtension(ext) != "image/jpeg" {
		exts, err := mime.ExtensionsByType("image/jpeg")
		if err != nil {
			panic(err)
		}
		for _, ext := range exts {
			names = append(names, name+ext)
		}
	}

	var dirsPath []string
	if filepath.IsAbs(name) {
		dirsPath = []string{""}
	} else {
		dirsPath = make([]string, len(dirs)+1)
		dirsPath[0] = "."
		copy(dirsPath[1:], dirs)
	}

	for _, dir := range dirsPath {
		for _, name := range names {
			path := filepath.Join(dir, name)

			stat, err := os.Stat(path)
			if err != nil {
				if os.IsNotExist(err) {
					continue
				}
				return "", err
			}

			if stat.IsDir() {
				continue
			}

			return path, nil
		}
	}

	return "", fmt.Errorf("%#v not found (path=%s)", name, strings.Join(dirsPath, ":"))
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
	return logrus.Level(*l).String()
}

func (l *logLevelValue) Type() string {
	return "string"
}
