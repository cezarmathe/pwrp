package recording

import (
	"github.com/cezarmathe/pwrp/gitops"
	log "github.com/sirupsen/logrus"
)

/*Recorder is a struct that does the recording process.*/
type Recorder struct {
	Config *Config

	logger *log.Logger
}

/*NewRecorder creates a new Recorder with the specified configuration.*/
func NewRecorder(config *Config, logger *log.Logger) *Recorder {
	return &Recorder{config, logger}
}

func (recorder *Recorder) checkIfShouldSkip(shouldSkip bool) bool {
	return recorder.Config.Skips.All || shouldSkip
}

/*Record starts the recording process.*/
func (recorder *Recorder) Record() bool {
	log.Debug("recorder.Record(): ", "called")
	log.Info("recording: ", "started")

	log.Debug("recorder.Record(): ", "starting iteration over repository list")
	for _, repositoryURL := range recorder.Config.Repositories {
		repository, err := gitops.Clone(repositoryURL, recorder.Config.StoragePath)
	}
	return true
}
