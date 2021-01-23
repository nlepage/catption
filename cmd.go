package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	version = "master"

	logLevel = logrus.InfoLevel

	cmd = &cobra.Command{
		Short: "Cat caption generator CLI",
		Args:  cobra.ExactArgs(1),
		PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
			logrus.SetLevel(logLevel)

			if err := initConfig(); err != nil {
				return err
			}

			return nil
		},
	}
)

func init() {
	cmd.Version = version

	cmd.PersistentFlags().Var((*logLevelValue)(&logLevel), "logLevel", "Log level")
}
