package recording

import (
	"net/url"
	"os"
	"pppi/git"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/phayes/permbits"
)

/*Config contains configurations for the recording process*/
type Config struct {
	Repositories []string     `toml:"repositories"`
	Protocol     git.Protocol `toml:"protocol"`
	StoragePath  string       `toml:"storage_path"`
	Skips        struct {
		MissingBranch      bool `toml:"missing_branch"`
		BadURL             bool `toml:"bad_url"`
		NoPermissions      bool `toml:"no_permissions"`
		BadProtocol        bool `toml:"bad_protocol"`
		MissingStoragePath bool `toml:"missing_storage_path"`
		All                bool `toml:"all"`
	} `toml:"skips"`
}

var (
	/*the actual configuration*/
	config Config

	/*the default path for saving*/
	defaultPath string
)

/*NewConfig creates a new dummy configuration file*/
func NewConfig() Config {
	logrus.Debug("creating new dummy recording config")
	return Config{
		Repositories: []string{},
		Protocol:     git.GIT,
		StoragePath:  "/home/username/.local/share/pppi-storage",
	}
}

/*SetConfig sets the configuration file for recording.*/
func SetConfig(cfg Config) {
	logrus.Debug("setting the recording config")
	config = cfg
}

/*ValidateConfig checks the configuration and validates its integrity*/
func ValidateConfig() bool {
	logrus.Info("validating the recording configuration")
	shouldContinue := true

	/*checking if any repositories were provided*/
	if len(config.Repositories) == 0 {
		checkConfigError(false, &shouldContinue, ErrNoRepositories)
		shouldContinue = false
	}

	/*checking if the protocol is good*/
	if config.Protocol == "" {
		config.Protocol = git.GIT
	}
	if !(config.Protocol == git.GIT || config.Protocol == git.HTTPS || config.Protocol == git.SSH) {
		checkConfigError(config.Skips.BadProtocol, &shouldContinue, NewErrBadProtocol(config.Protocol))
	}

	/*checking if the storage path is valid*/
	if fileMode, pathErr := os.Stat(config.StoragePath); pathErr != nil {
		/*check if the directory exists and create it if it doesn't*/
		if os.IsNotExist(pathErr) {
			logrus.Info("provided recording storage path does not exist")
			var permissions permbits.PermissionBits
			permissions = 744
			fileMode := new(os.FileMode)
			permbits.UpdateFileMode(fileMode, permissions)
			err := os.Mkdir(config.StoragePath, *fileMode)
			if err != nil {
				checkConfigError(false, &shouldContinue, NewErrCreateStorageDir(config.StoragePath))
			} else {
				logrus.Info("created recording storage path")
			}
		} else { /*another error*/
			checkConfigError(false, &shouldContinue, pathErr)
		}
	} else {
		permissions := permbits.FileMode(fileMode.Mode())
		/*check if the directory has the proper permissions*/
		if !permissions.UserExecute() || !permissions.UserRead() || !permissions.UserWrite() {
			checkConfigError(config.Skips.NoPermissions, &shouldContinue, NewErrNoPermissions(config.StoragePath))
		}
	}

	/*checking each repository's URL and protocol*/
	logrus.Debug("validating repository URL's")
	for index := range config.Repositories {
		logrus.Debug("checking ", config.Repositories[index])
		// todo 14/03/2019: split url string by "://" and check the protocol
		/*check if it has a protocol and add it if it's missing*/
		if !strings.HasPrefix(config.Repositories[index], string(git.GIT)) || !strings.HasPrefix(config.Repositories[index], string(git.SSH)) || !strings.HasPrefix(config.Repositories[index], string(git.HTTPS)) {
			config.Repositories[index] = strings.Join([]string{
				string(config.Protocol),
				"://",
				config.Repositories[index],
			}, "")
		}

		/*checking if the url is valid*/
		_, err := url.ParseRequestURI(config.Repositories[index])
		if err != nil {
			checkConfigError(config.Skips.BadURL, &shouldContinue, NewErrBadURL(config.Repositories[index]))
		}
	}
	if shouldContinue {
		logrus.Info("validated the recording configuration")
	}
	return shouldContinue
}

/*checkConfigError checks a configuration error and, based on if it should be skipped or not,
the error is logged to Warn or Error and the shouldContinue flag is set to false*/
func checkConfigError(shouldSkip bool, shouldContinue *bool, text error) {
	if shouldSkip {
		logrus.Warn(text)
		return
	}
	*shouldContinue = false
	logrus.Error(text)
}
