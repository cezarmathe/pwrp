package config

import (
	"io/ioutil"
	"os"
	"pppi/recording"

	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
)

/*Config is a container for the entire utility configuration*/
type Config struct {
	Recording recording.Config `toml:"recording"`
}

func init() {
	/*check whether to use debug level or not*/
	logLevel := os.Getenv("DEBUG_LOG_LEVEL")
	if logLevel != "" {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
}

/*LoadConfig loads the configuration from a file and returns it*/
func LoadConfig(filename string) (Config, error) {
	logrus.Debug("loading the configuration")
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}
	cfg := new(Config)
	err = toml.Unmarshal(file, cfg)
	logrus.Info("loaded the configuration")
	return *cfg, err
}
