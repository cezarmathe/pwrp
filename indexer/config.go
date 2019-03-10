package indexer

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

// Protocol used for cloning
type Protocol string

// Cloning protocols
const (
	HTTPS           Protocol = "https"
	SSH             Protocol = "ssh"
	GIT             Protocol = "git"
	DefaultProtocol Protocol = GIT
)

// Destination directory for indexing
var (
	DefaultDir string
)

// The configuration object
var config Config

// Config is a configuration data structure for indexing projects
type Config struct {
	Repositories         []string `toml:"repositories"`
	CloneProtocol        Protocol `toml:"protocol, omitempty"`
	DestinationDirectory string   `toml:"dest_dir, omitempty"`
	IgnoreFailed         bool     `toml:"ignore_failed, omitempty"`
}

func init() {
	logrus.SetLevel(logrus.DebugLevel)

	homedir, err := os.Getwd()
	if err != nil {
		logrus.Warn("could not open homedir")
	}
	DefaultDir = homedir + "/.pppi_index_dir"
}

func LoadConfig(filename string) bool {
	logrus.Debug("loading indexing configuration file")

	configFile, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.Error(err)
		return false
	}

	err = toml.Unmarshal(configFile, &config)
	if err != nil {
		logrus.Error(err)
		return false
	}

	logrus.Info("loaded indexing configuration file")

	return true
}
