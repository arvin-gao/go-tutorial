package packages

import (
	"os"

	"github.com/sirupsen/logrus"
)

func newLogger(f *os.File) *logrus.Logger {
	l := logrus.New()
	if f != nil {
		l.SetOutput(f)
	}
	l.SetFormatter(&logrus.TextFormatter{})
	l.Level = logrus.TraceLevel
	return l
}
