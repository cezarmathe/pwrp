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
	"github.com/spf13/viper"

	cfg "github.com/cezarmathe/pwrp/config"
)

var (
	configCmd = &cobra.Command{
		Use:   "config",
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

	log.Trace("running config command")

	log.Debug("initializing config logging")
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
		runConfigValidation(true)
	}
}

func runConfigValidation(mustRewrite bool) bool {
	log.DebugFunctionCalled()
	defer log.DebugFunctionReturned()

	log.Info("validating the configuration file")

	log.Debug("initialize config logging")
	cfg.InitLogging(log.GetParams())

	log.Trace("running the config validation")
	pass := cfg.ValidateConfig(config)

	log.Info("validation passed: ", pass)

	if pass {
		log.Trace("rewriting configuration file")

		err := config.WriteConfig()
		if err != nil {
			if mustRewrite {
				log.ErrorErr(err, "failed to rewrite the configuration file")
			} else {
				log.WarnErr(err, "failed to rewrite the configuration file")
			}
		} else {
			log.Info("rewrote the configuration file")
		}
		return true
	}
	return false
}

func runConfigExport() {
	log.DebugFunctionCalled()
	defer log.DebugFunctionReturned()

	/*create a new dummy configuration*/
	dummyConfig := viper.New()
	cfg.NewDummyConfig(dummyConfig)

	/*write the configuration to a file*/
	data, _ := toml.Marshal(dummyConfig)
	err := ioutil.WriteFile(configFilePath, data, 0744)
	if err != nil {
		log.ErrorErr(err, "failed to export dummy configuration to ", configFilePath)
	} else {
		log.Info("exported dummy configuration")
	}
}
