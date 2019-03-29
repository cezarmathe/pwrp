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

	"github.com/cezarmathe/pwrp/smartlogger"
	"github.com/phayes/permbits"

	"github.com/cezarmathe/pwrp/recording"
)

var (
	log *smartlogger.SmartLogger
)

/*Config is a container for the entire utility configuration*/
type Config struct {
	Recording   *recording.Config `toml:"recording"`
	StoragePath string            `toml:"storage_path"`
}

/*ValidateConfig validates the configuration*/
func ValidateConfig(config *Config) bool {
	log.Debug("called")

	shouldContinue := true

	/*checking if the storage path is valid and has proper permissions*/
	log.Debug("storage path validation: ", "checking")
	if _, pathErr := os.Stat(config.StoragePath); pathErr != nil {
		/*check if the directory exists and create it if it doesn't*/
		log.Trace("ValidateConfig(): ", "checking if path exists")
		if os.IsNotExist(pathErr) {
			log.Info("storage path validation: ", "storage path does not exist, attempting to create it")
			var permissions permbits.PermissionBits
			permissions = 0
			permissions += permbits.UserRead + permbits.UserWrite + permbits.UserExecute
			permissions += permbits.GroupRead + permbits.GroupExecute
			permissions += permbits.OtherRead
			fileMode := new(os.FileMode)
			permbits.UpdateFileMode(fileMode, permissions)
			log.Trace("ValidateConfig(): ", "creating storage directory")
			err := os.Mkdir(config.StoragePath, *fileMode)
			if err != nil {
				log.Debug("storage path validation: ", "failed to create storage directory")
				log.Error("storage path validation: ", NewErrCreateStorageDir(config.StoragePath))
				shouldContinue = false
			} else {
				log.Info("storage path validation: ", "created storage directory at "+config.StoragePath)
			}
		} else { /*another error*/
			log.Error("storage path validation: ", "unknown error - ", pathErr)
			shouldContinue = false
		}
	}
	if permissions, err := permbits.Stat(config.StoragePath); err == nil {
		/*check if the directory has the proper permissions*/
		log.Debug("storage path validation: ", "checking directory permissions")
		if !permissions.UserExecute() || !permissions.UserRead() || !permissions.UserWrite() {
			log.Error("storage path validation: ", NewErrNoPermissions(config.StoragePath))
			shouldContinue = false
		}
	} else {
		log.Error("storage path validation: ", "unknown error - ", err)
	}

	shouldContinue = recording.NewRecorder(config.Recording, nil).ValidateConfig() && shouldContinue
	log.Trace("ValidateConfig(): ", "returned")
	return shouldContinue
}
