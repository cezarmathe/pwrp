package recording

import (
	"net/url"

	"strings"

	"github.com/cezarmathe/pwrp/git"
	log "github.com/sirupsen/logrus"
)

/*Config contains configurations for the recording process*/
type Config struct {
	Repositories []string     `toml:"repositories"`
	Protocol     git.Protocol `toml:"protocol"`
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
		Protocol:     git.GIT,
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
		recorder.Config.Protocol = git.GIT
	}
	if !(recorder.Config.Protocol == git.GIT || recorder.Config.Protocol == git.HTTPS || recorder.Config.Protocol == git.SSH) {
		log.Trace("recorder.ValidateConfig(): ", "protocol is bad")
		checkConfigError(recorder.Config.Skips.BadProtocol, &shouldContinue, NewErrBadProtocol(recorder.Config.Protocol))
	}

	/*checking each repository's URL and protocol*/
	log.Debug("validating repository URL's")
	for index := range recorder.Config.Repositories {
		log.Debug("checking ", recorder.Config.Repositories[index])
		// todo 14/03/2019: split url string by "://" and check the protocol
		/*check if it has a protocol and add it if it's missing*/
		if !strings.HasPrefix(recorder.Config.Repositories[index], string(git.GIT)) || !strings.HasPrefix(recorder.Config.Repositories[index], string(git.SSH)) || !strings.HasPrefix(recorder.Config.Repositories[index], string(git.HTTPS)) {
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
