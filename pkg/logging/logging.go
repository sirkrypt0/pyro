package logging

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log = &logrus.Logger{
	Out: os.Stderr,
	Formatter: &logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	},
	Hooks: make(logrus.LevelHooks),
	Level: logrus.DebugLevel,
}

func GetLogger(pkg string) *logrus.Entry {
	return log.WithField("pkg", pkg)
}
