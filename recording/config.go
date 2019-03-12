package recording

import "pppi/git"

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
	return Config{
		Repositories: []string{},
		Protocol:     git.GIT,
		StoragePath:  "/home/username/.local/share/pppi-storage",
	}
}

/*SetConfig sets the configuration file for recording.*/
func SetConfig(cfg Config) {
	config = cfg
}

/*validateConfig checks the configuration and validates its integrity*/
func validateConfig() []error {
	errs := make([]error, 5)

	return errs
}
