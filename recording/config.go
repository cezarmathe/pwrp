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

/*validateConfig checks the configuration and validates its integrity*/
func validateConfig() []error {
	logrus.Info("validating config file")
	errs := make([]error, 5)

	/*checking if any repositories were provided*/
	if len(config.Repositories) == 0 {
		logrus.Debug("empty repository url list")
		errs = append(errs, ErrNoRepositories)
	}

	/*checking if the protocol is good*/
	// todo assign the default protocol if it is missing
	if !(config.Protocol == git.GIT || config.Protocol == git.HTTPS || config.Protocol == git.SSH) {
		logrus.Debug("provided protocol is bad")
		errs = append(errs, NewErrBadProtocol(config.Protocol))
	}

	/*checking if the storage path is valid*/
	if fileMode, pathErr := os.Stat(config.StoragePath); pathErr != nil {
		/*check if the directory exists and create it if it doesn't*/
		if os.IsNotExist(pathErr) {
			logrus.Debug("provided recording storage path does not exist")
			var permissions permbits.PermissionBits
			permissions = 744
			fileMode := new(os.FileMode)
			permbits.UpdateFileMode(fileMode, permissions)
			err := os.Mkdir(config.StoragePath, *fileMode)
			if err != nil {
				logrus.Debug("encountered an error when creating the storage directory")
				errs = append(errs, err)
			} else {
				logrus.Info("created recording storage path")
			}
		} else { /*another error*/
			logrus.Debug("another recording storage path error")
			errs = append(errs, pathErr)
		}
	} else {
		permissions := permbits.FileMode(fileMode.Mode())
		/*check if the directory has the proper permissions*/
		if !permissions.UserExecute() || !permissions.UserRead() || !permissions.UserWrite() {
			logrus.Debug("bad storage path directory permissions")
			errs = append(errs, NewErrNoPermissions(config.StoragePath))
		}
	}

	/*checking each repository's URL and protocol*/
	logrus.Debug("validating repository URL's")
	for index, repositoryURL := range config.Repositories {
		logrus.Debug("checking", repositoryURL)

		/*check if it has a protocol*/
		if !strings.HasPrefix(repositoryURL, string(git.GIT)) || !strings.HasPrefix(repositoryURL, string(git.SSH)) || !strings.HasPrefix(repositoryURL, string(git.HTTPS)) {
			config.Repositories[index] = strings.Join([]string{}, "")
			// todo finish
		}

		/*checking if the url is valid*/
		_, err := url.ParseRequestURI(repositoryURL)
		if err != nil {
			errs = append(errs, NewErrBadURL(repositoryURL))
		}
	}

	return errs
}
