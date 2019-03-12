package recording

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	/*check whether to use debug level or not*/
	logLevel := os.Getenv("DEBUG_LOG_LEVEL")
	if logLevel != "" {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	homedir, err := os.Getwd()
	if err != nil {
		logrus.Warn("could not open homedir")
	}
	defaultPath = homedir + "/.local/share/pppi-storage"
}
