package main

import (
	"flag"
	"os"
	"pppi/config"
	"pppi/recording"

	"github.com/sirupsen/logrus"
)

var (
	cfg                    config.Config
	indexingConfigFileName string
)

func init() {
	/*check whether to use debug level or not*/
	logLevel := os.Getenv("DEBUG_LOG_LEVEL")
	if logLevel != "" {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	/*parse flags*/
	logrus.Debug("parsing command line flags")

	flag.StringVar(&indexingConfigFileName, "cfg", "", "file path for the configuration file")

	flag.Parse()

	if indexingConfigFileName == "" {
		logrus.Fatal("no indexing config file provided")
	}
	logrus.Debug("parsed command line flags")

	/*load configuration*/
	cfg, err := config.LoadConfig(indexingConfigFileName)
	if err != nil {
		logrus.Fatal("encountered an error when reading the configuration ", err)
	}

	/*set appropriate configurations*/
	recording.SetConfig(cfg.Recording)
}

func main() {
	/*validate recording config*/
	recordingErrs := recording.ValidateConfig()
	if len(recordingErrs) != 0 {
		for _, err := range recordingErrs {
			if err != nil {
				logrus.Warn(err)
			}
		}
	}
}
