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

	"github.com/cezarmathe/pwrp/config/keys"
	"github.com/spf13/viper"

	"github.com/cezarmathe/pwrp/gitops"
)

/*Config contains configurations for the recording process*/
type Config struct {
	Repositories []string        `toml:"repositories"`
	Protocol     gitops.Protocol `toml:"protocol"`
	Skips        struct {
		MissingBranch bool `toml:"missing_branch"`
		BadURL        bool `toml:"bad_url"`
		BadProtocol   bool `toml:"bad_protocol"`
		All           bool `toml:"all"`
	} `toml:"skips"`
}

/*NewDummyConfig creates a new dummy configuration file*/
func NewDummyConfig(dummyConfig *viper.Viper) {
	log.DebugFunctionCalled()

	dummyConfig.Set(keys.RecordingRepositoryListKey, []string{})
	dummyConfig.Set(keys.RecordingProtocolKey, gitops.DefaultProtocol)

	dummyConfig.Set(keys.RecordingSkipsMissingBranchKey, false)
	dummyConfig.Set(keys.RecordingSkipsBadURLKey, false)
	dummyConfig.Set(keys.RecordingSkipsBadProtocolKey, false)
	dummyConfig.Set(keys.RecordingSkipsAllKey, false)

	log.DebugFunctionReturned()
}

/*ValidateConfig checks the configuration and validates its integrity*/
func (recorder *Recorder) ValidateConfig() bool {
	log.DebugFunctionCalled()
	log.Info("started recording config validation")

	configIsValid := true

	/*checking if any repositories were provided*/
	log.Trace("checking if any repositories were provided")
	if len(recorder.Config.GetStringSlice(keys.RecordingRepositoryListKey)) == 0 {
		checkConfigError(false, &configIsValid, ErrNoRepositories)
		configIsValid = false
	} else {
		log.Trace("repository list is not empty")
	}

	/*checking if the protocol is good*/
	log.Trace("checking cloning protocol")
	if recorder.Config.Get(keys.RecordingProtocolKey) == "" {
		log.Trace("using the default protocol")
		recorder.Config.Set(keys.RecordingProtocolKey, gitops.DefaultProtocol)
	}
	if !(gitops.NewProtocol(recorder.Config.GetString(keys.RecordingProtocolKey)) == gitops.GIT ||
		gitops.NewProtocol(recorder.Config.GetString(keys.RecordingProtocolKey)) == gitops.HTTPS ||
		gitops.NewProtocol(recorder.Config.GetString(keys.RecordingProtocolKey)) == gitops.SSH) {

		checkConfigError(
			recorder.Config.GetBool(keys.RecordingSkipsBadProtocolKey),
			&configIsValid,
			NewErrBadProtocol(gitops.NewProtocol(recorder.Config.GetString(keys.RecordingProtocolKey))))
	}

	/*checking each repository's URL and protocol*/
	log.Trace("checking repository URL's")
	repositories := recorder.Config.GetStringSlice(keys.RecordingRepositoryListKey)
	for index := range repositories {
		repoUrl := &repositories[index]
		log.Trace("checking URL ", *repoUrl)

		/*check if the URL has a protocol*/
		splitUrl := strings.Split(*repoUrl, "://")
		if len(splitUrl) == 1 {
			log.Trace("URL ", *repoUrl, " has no protocol, adding the default protocol ", gitops.DefaultProtocol)
			*repoUrl = strings.Join([]string{string(gitops.DefaultProtocol), "://", *repoUrl}, "")

		} else if urlHasProtocol(*repoUrl, gitops.GIT) || urlHasProtocol(*repoUrl, gitops.SSH) || urlHasProtocol(*repoUrl, gitops.HTTPS) {
			log.Trace("the url ", *repoUrl, " has a valid protocol")
		} else {
			checkConfigError(recorder.Config.GetBool(keys.RecordingSkipsBadProtocolKey), &configIsValid, NewErrBadURL(*repoUrl))
		}

		/*checking if the repoUrl is valid*/
		_, err := url.ParseRequestURI(*repoUrl)
		if err != nil {
			checkConfigError(recorder.Config.GetBool(keys.RecordingSkipsBadURLKey), &configIsValid, NewErrBadURL(*repoUrl))
		}
	}
	recorder.Config.Set(keys.RecordingRepositoryListKey, repositories)
	if configIsValid {
		log.Info("validated the recording configuration")
	} else {
		log.Warn("the recording configuration is invalid")
	}
	return configIsValid
}

/*checkConfigError checks a configuration error and, based on if it should be skipped or not,
the error is logged to Warn or Error and the shouldContinue flag is set to false.*/
func checkConfigError(shouldSkip bool, configIsValid *bool, err error) {
	if shouldSkip {
		log.WarnErr(err)
		return
	}
	*configIsValid = false
	log.ErrorErr(err)
}

/*urlHasProtocol checks if the given URL has the given protocol*/
func urlHasProtocol(url string, protocol gitops.Protocol) bool {
	return strings.HasPrefix(url, string(protocol))
}
