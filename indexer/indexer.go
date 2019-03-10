package indexer

import "github.com/sirupsen/logrus"

func StartJob() (error, bool) {
	logrus.Info("starting indexing job")
	if config.DestinationDirectory == "" && DefaultDir == "" {
		logrus.Error("no directory available for working")
		return NoDirAvailable, false
	}

	var baseDirPath string

	if config.DestinationDirectory != "" {
		logrus.Info("using provided work directory from the config file")
		baseDirPath = config.DestinationDirectory
	} else {
		logrus.Info("using the default working directory($(pwd)/.pppi_index_dir)")
		baseDirPath = DefaultDir
	}

	baseDirPath = baseDirPath
	return nil, true
}
