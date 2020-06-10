package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

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
