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
		if err != nil {
			log.Fatal("recording: ", "fatal error encountered when cloning - ", err)
			return false
		}
		log.Info("recording: ", "repository "+repositoryURL+"cloned successfully")

		_, err = repository.Branch("_pwrp")

	}
	return true
}
