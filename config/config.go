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

package config

import (
	"os"

	"github.com/cezarmathe/pwrp/config/keys"
	"github.com/mitchellh/go-homedir"
	"github.com/phayes/permbits"
	"github.com/spf13/viper"

	"github.com/cezarmathe/pwrp/recording"
)

/*Config is a container for the entire utility configuration*/
type Config struct {
	Recording   *recording.Config `toml:"recording"`
	StoragePath string            `toml:"storage_path"`
}

func NewDummyConfig(dummyConfig *viper.Viper) {
	log.DebugFunctionCalled(*dummyConfig)

	log.Debug("initialize recording logging")
	recording.InitLogging(log.GetParams())

	/*find home directory.*/
	log.Trace("finding home directory")
	home, err := homedir.Dir()
	if err != nil {
		log.FatalErr(err, "encountered an error when trying to find the home directory")
	}

	recording.NewDummyConfig(dummyConfig)
	viper.Set(keys.StoragePathKey, home+"/.local/share/pwrp")

	log.DebugFunctionReturned()
}

/*ValidateConfig validates the configuration*/
func ValidateConfig(globalConfig *viper.Viper) bool {
	log.DebugFunctionCalled(globalConfig)

	var shouldContinue = true

	log.Trace("storage path: ", globalConfig.GetString(keys.StoragePathKey))

	/*checking if the storage path is valid and has proper permissions*/
	log.Trace("checking if the storage path is valid and it exists")
	if _, pathErr := os.Stat(globalConfig.GetString(keys.StoragePathKey)); pathErr != nil {

		/*check if the directory exists and create it if it doesn't*/
		log.Trace("checking if the path exists")
		if os.IsNotExist(pathErr) {
			log.Warn("storage path ,", globalConfig.GetString(keys.StoragePathKey), " does not exist, attempting to create it")

			var permissions permbits.PermissionBits

			permissions = 0
			permissions += permbits.UserRead + permbits.UserWrite + permbits.UserExecute
			permissions += permbits.GroupRead + permbits.GroupExecute
			permissions += permbits.OtherRead

			fileMode := new(os.FileMode)
			permbits.UpdateFileMode(fileMode, permissions)

			log.Trace("creating storage directory")
			err := os.MkdirAll(globalConfig.GetString(keys.StoragePathKey), *fileMode)
			if err != nil {
				log.ErrorErr(NewErrCreateStorageDir(globalConfig.GetString(keys.StoragePathKey)), "storage path validation failed")
				shouldContinue = false
			} else {
				log.Info("created storage directory at " + globalConfig.GetString(keys.StoragePathKey))
			}
		} else { /*another error*/
			log.ErrorErr(pathErr, "unknown error")
			shouldContinue = false
		}
	} else {
		log.Trace("storage path exists")
	}

	/*check if the directory has the proper permissions*/
	log.Trace("checking if the storage directory has the proper permissions")
	if permissions, err := permbits.Stat(globalConfig.GetString(keys.StoragePathKey)); err == nil {
		if !permissions.UserExecute() || !permissions.UserRead() || !permissions.UserWrite() {
			log.ErrorErr(NewErrNoPermissions(globalConfig.GetString(keys.StoragePathKey)))
			shouldContinue = false
		}
	} else {
		log.ErrorErr(err, "unknown error")
	}

	/*running the recording validation*/
	log.Trace("running the recorder validation")
	recording.InitLogging(log.GetParams())
	shouldContinue = recording.NewRecorder(globalConfig).ValidateConfig() && shouldContinue

	log.DebugFunctionReturned(shouldContinue)
	return shouldContinue
}
