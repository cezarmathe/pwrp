package main

import (
	"flag"
	"pppi/indexer"

	"github.com/sirupsen/logrus"
)

var (
	indexingConfigFileName string
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)

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
