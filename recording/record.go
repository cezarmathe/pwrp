/*
 * PWRP - Personal Work Recorder Processor
 * Copyright (C) 2019  Cezar Mathe <cezarmathe@gmail.com> [https://cezarmathe.com]
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published
 * by the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package recording

import (
	"github.com/spf13/viper"

	"github.com/cezarmathe/pwrp/config/keys"

	"github.com/cezarmathe/pwrp/gitops"
)

const (
	MetadataBranchName = "_pwrp-metadata"
)

/*Recorder is a struct that does the recording process.*/
type Recorder struct {
	Config *viper.Viper
}

/*NewRecorder creates a new Recorder with the specified configuration.*/
func NewRecorder(config *viper.Viper) *Recorder {
	return &Recorder{config}
}

func (recorder *Recorder) checkIfShouldSkip(shouldSkip bool) bool {
	return recorder.Config.GetBool(keys.RecordingSkipsAllKey) || shouldSkip
}

/*Record starts the recording process.*/
func (recorder *Recorder) Record() bool {
	log.DebugFunctionCalled()
	log.Info("recording process started")

	var shouldContinue = true

	log.Debug("initializing gitops logging")
	gitops.InitLogging(log.GetParams())

	log.Trace("storage path: ", recorder.Config.Get(keys.StoragePathKey))

	log.Trace("iterating over repository list")
	for _, repositoryURL := range recorder.Config.GetStringSlice(keys.RecordingRepositoryListKey) {
		log.Trace("operating on URL ", repositoryURL)
		repository, err := gitops.Clone(repositoryURL, recorder.Config.GetString(keys.StoragePathKey))
		if err != nil {
			log.ErrorErr(err, "error encountered when loading the repository ", repositoryURL)
			shouldContinue = false
			break
		}
		log.Info("repository ", repositoryURL, " loaded successfully")

		_, _ = repository.Branch(MetadataBranchName)

	}
	log.Trace("finished iterating over repository list")
	log.Info("recording process finished")
	log.DebugFunctionReturned(shouldContinue)
	return shouldContinue
}
