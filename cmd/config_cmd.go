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

package cmd

import (
	"io/ioutil"

	"github.com/pelletier/go-toml"
	"github.com/spf13/cobra"

	cfg "github.com/cezarmathe/pwrp/config"
)

var (
	configCmd = &cobra.Command{
		Use:   "_config",
		Short: "Configuration related tools",
		Run:   runConfigCmd,
	}
)

var (
	configCmdExportFlag   = false
	configCmdValidateFlag = false
)

func init() {
	configCmd.Flags().BoolVarP(&configCmdExportFlag, "export", "e", false, "export a dummy configuration")
	configCmd.Flags().BoolVar(&configCmdValidateFlag, "validate", false, "validate the configuration file")
}

func runConfigCmd(cmd *cobra.Command, args []string) {
	log.DebugFunctionCalled(*cmd, args)
	defer log.DebugFunctionReturned()

	log.Trace("running _config command")

	log.Debug("initializing _config logging")
	cfg.InitLogging(log.GetParams())

	if !configCmdValidateFlag && !configCmdExportFlag {
		log.Debug("both flags are empty; activating validate")
		configCmdValidateFlag = true
	}

	if configCmdExportFlag {
		log.Trace("exporting dummy configuration file")
		runConfigExport()
	}
	if configCmdValidateFlag {
		log.Info("validating the configuration file")
		log.Info("validation passed: ", runConfigValidation())
	}
}

func runConfigValidation() bool {
	log.DebugFunctionCalled()

	log.Info("validating the configuration file")

	log.Debug("initialize _config logging")
	cfg.InitLogging(log.GetParams())

	log.Trace("running the _config validation")
	pass := cfg.ValidateConfig(_config)

	log.DebugFunctionReturned(pass)
	return pass
}

func runConfigExport() {
	log.DebugFunctionCalled()
	defer log.DebugFunctionReturned()

	/*create a new dummy configuration*/
	dummyConfig := cfg.NewDummyConfig()

	/*write the configuration to a file*/
	data, _ := toml.Marshal(dummyConfig)
	err := ioutil.WriteFile(configFilePath, data, 0744)
	if err != nil {
		log.ErrorErr(err, "failed to export dummy configuration to ", configFilePath)
	} else {
		log.Info("exported dummy configuration")
	}
}
