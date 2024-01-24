package implement

import (
	"log"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func logrusGuides() {
	l := logrus.New()
	// Set formatter. includes log.JSONFormatter{} and log.TextFormatter{}.
	l.SetFormatter(&logrus.TextFormatter{
		DisableColors:    true,
		DisableTimestamp: true,
	})

	// Will log anything that is info or above (warn, error, fatal, panic). Default.
	logrus.SetLevel(logrus.InfoLevel)

	// SetReportCaller If you wish to add the calling method as a field, instruct the logger via:
	// e.g. field: "method":"github.com/sirupsen/arcticcreatures.migrate"
	logrus.SetReportCaller(true)

	// Custom fields.
	l.WithField("my key", "my value").Info("this is info message")
	l.WithFields(logrus.Fields{
		"my key1": "value",
		"my key2": "value",
	}).Info("this is info message")
	logWithKey := logrus.WithField("key", "value")
	logWithKey.Error("this is error message1")
	logWithKey.Error("this is error message2")

	// Set output source.
	l.Out = os.Stdout
}

func logrusWithFile() {
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		logrus.Fatal(err)
	}
	l := logrus.New()
	l.Out = file
	l.Info("info")
}

func logrusAsWriter() {
	l := logrus.New()
	w := l.Writer()
	defer w.Close()

	_ = http.Server{
		// create a stdlib log.Logger that writes to logrus.Logger.
		ErrorLog: log.New(w, "", 0),
	}
}
