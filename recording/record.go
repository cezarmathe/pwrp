package recording

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)

	homedir, err := os.Getwd()
	if err != nil {
		logrus.Warn("could not open homedir")
	}
	defaultPath = homedir + "/.local/share/pppi-storage"
}
