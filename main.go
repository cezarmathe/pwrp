package main

import (
	"flag"
	"os"
	"pppi/indexer"

	"github.com/sirupsen/logrus"
)

var (
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

	logrus.Debug("parsing command line flags")

	flag.StringVar(&indexingConfigFileName, "i", "", "file path for the indexing configuration file")

	flag.Parse()

	if indexingConfigFileName == "" {
		logrus.Fatal("no indexing config file provided")
	}
	logrus.Debug("parsed command line flags")
}

func main() {
	indexer.LoadConfig(indexingConfigFileName)

	err, canContinue := indexer.StartJob()
	if !canContinue {
		logrus.Fatal(err)
	}
	if err != nil {
		logrus.Error(err)
	}
}
