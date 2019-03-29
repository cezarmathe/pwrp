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

	"github.com/phayes/permbits"

	"github.com/cezarmathe/pwrp/recording"
)

/*Config is a container for the entire utility configuration*/
type Config struct {
	Recording   *recording.Config `toml:"recording"`
	StoragePath string            `toml:"storage_path"`
}

/*ValidateConfig validates the configuration*/
func ValidateConfig(config *Config) bool {
	log.DebugFunctionCalled(config)

	var shouldContinue = true

	/*checking if the storage path is valid and has proper permissions*/
	log.Trace("checking if the storage path is valid and has proper permissions")
	if _, pathErr := os.Stat(config.StoragePath); pathErr != nil {
		log.WarnErr(pathErr)

		/*check if the directory exists and create it if it doesn't*/
		log.Trace("checking if the path exists")
		if os.IsNotExist(pathErr) {
			log.Warn("storage path does not exist, attempting to create it")

			var permissions permbits.PermissionBits

			permissions = 0
			permissions += permbits.UserRead + permbits.UserWrite + permbits.UserExecute
			permissions += permbits.GroupRead + permbits.GroupExecute
			permissions += permbits.OtherRead

			fileMode := new(os.FileMode)
			permbits.UpdateFileMode(fileMode, permissions)

			log.Trace("creating storage directory")
			err := os.MkdirAll(config.StoragePath, *fileMode)
			if err != nil {
				log.ErrorErr(NewErrCreateStorageDir(config.StoragePath), "storage path validation failed")
				shouldContinue = false
			} else {
				log.Info("created storage directory at " + config.StoragePath)
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
	if permissions, err := permbits.Stat(config.StoragePath); err == nil {
		if !permissions.UserExecute() || !permissions.UserRead() || !permissions.UserWrite() {
			log.ErrorErr(NewErrNoPermissions(config.StoragePath))
			shouldContinue = false
		}
	} else {
		log.ErrorErr(err, "unknown error")
	}

	/*running the recording validation*/
	log.Trace("running the recorder validation")
	recording.InitLogging(log.GetParams())
	shouldContinue = recording.NewRecorder(config.Recording).ValidateConfig() && shouldContinue

	log.DebugFunctionReturned(shouldContinue)
	return shouldContinue
}
