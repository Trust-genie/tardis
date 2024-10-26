package logger

import "github.com/sirupsen/logrus"

func init() {
	Log = logrus.New()
}

var Log *logrus.Logger
