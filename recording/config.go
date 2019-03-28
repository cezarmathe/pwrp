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
	"net/url"

	"strings"

	"github.com/cezarmathe/pwrp/gitops"
	log "github.com/sirupsen/logrus"
)

/*Config contains configurations for the recording process*/
type Config struct {
	Repositories []string        `toml:"repositories"`
	Protocol     gitops.Protocol `toml:"protocol"`
	StoragePath  string
	Skips        struct {
		MissingBranch bool `toml:"missing_branch"`
		BadURL        bool `toml:"bad_url"`
		BadProtocol   bool `toml:"bad_protocol"`
		All           bool `toml:"all"`
	} `toml:"skips"`
}

/*NewDummyConfig creates a new dummy configuration file*/
func NewDummyConfig() *Config {
	log.Debug("creating new dummy recording config")
	return &Config{
		Repositories: []string{},
		Protocol:     gitops.GIT,
		StoragePath:  "/home/username/.local/share/pwrp-storage",
	}
}

/*ValidateConfig checks the configuration and validates its integrity*/
func (recorder *Recorder) ValidateConfig() bool {
	log.Trace("recorder.ValidateConfig(): ", "called with ", recorder.Config)
	log.Info("recording config validation: ", "started")

	shouldContinue := true

	/*checking if any repositories were provided*/
	log.Debug("repository list validation: ", "checking")
	if len(recorder.Config.Repositories) == 0 {
		log.Trace("recorder.ValidateConfig(): ", "repository list is empty")
		checkConfigError(false, &shouldContinue, ErrNoRepositories)
		shouldContinue = false
	}
	log.Trace("recorder.ValidateConfig(): ", "repository list is not empty")

	/*checking if the protocol is good*/
	log.Debug("clone protocol validation: ", "checking")
	if recorder.Config.Protocol == "" {
		log.Debug("clone protocol validation: ", "setting default protocol")
		recorder.Config.Protocol = gitops.GIT
	}
	if !(recorder.Config.Protocol == gitops.GIT || recorder.Config.Protocol == gitops.HTTPS || recorder.Config.Protocol == gitops.SSH) {
		log.Trace("recorder.ValidateConfig(): ", "protocol is bad")
		checkConfigError(recorder.Config.Skips.BadProtocol, &shouldContinue, NewErrBadProtocol(recorder.Config.Protocol))
	}

	/*checking each repository's URL and protocol*/
	log.Debug("validating repository URL's")
	for index := range recorder.Config.Repositories {
		log.Debug("checking ", recorder.Config.Repositories[index])
		// todo 14/03/2019: split url string by "://" and check the protocol
		/*check if it has a protocol and add it if it's missing*/
		if !strings.HasPrefix(recorder.Config.Repositories[index], string(gitops.GIT)) || !strings.HasPrefix(recorder.Config.Repositories[index], string(gitops.SSH)) || !strings.HasPrefix(recorder.Config.Repositories[index], string(gitops.HTTPS)) {
			recorder.Config.Repositories[index] = strings.Join([]string{
				string(recorder.Config.Protocol),
				"://",
				recorder.Config.Repositories[index],
			}, "")
		}

		/*checking if the url is valid*/
		_, err := url.ParseRequestURI(recorder.Config.Repositories[index])
		if err != nil {
			checkConfigError(recorder.Config.Skips.BadURL, &shouldContinue, NewErrBadURL(recorder.Config.Repositories[index]))
		}
	}
	if shouldContinue {
		log.Info("validated the recording configuration")
	}
	return shouldContinue
}

/*checkConfigError checks a configuration error and, based on if it should be skipped or not,
the error is logged to Warn or Error and the shouldContinue flag is set to false.*/
func checkConfigError(shouldSkip bool, shouldContinue *bool, text error) {
	if shouldSkip {
		log.Warn(text)
		return
	}
	*shouldContinue = false
	log.Error(text)
}
