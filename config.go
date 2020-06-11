// +build !js,!wasm

package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.BindPFlag("dirs", cmd.Flags().Lookup("dir"))
}

func initConfig() error {
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
		logrus.Debug("No config file found")
	} else {
		logrus.Debugf("Using config file %s", viper.ConfigFileUsed())
	}

	dirs = viper.GetStringSlice("dirs")

	return nil
}
