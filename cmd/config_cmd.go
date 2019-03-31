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
	"github.com/spf13/cobra"

	cfg "github.com/cezarmathe/pwrp/config"
)

var (
	configCmd = &cobra.Command{
		Use:   "validate-config",
		Short: "Validate the configuration file",
		Run:   runConfigCmd,
	}
)

func runConfigCmd(cmd *cobra.Command, args []string) {
	log.DebugFunctionCalled(*cmd, args)
	defer log.DebugFunctionReturned()

	log.Info("running the validate config command")
	runConfigValidation()
}

func runConfigValidation() bool {
	log.DebugFunctionCalled()

	log.Debug("initialize config logging")
	cfg.InitLogging(log.GetParams())

	log.Trace("running the config validation")
	pass := cfg.ValidateConfig(config)
	log.DebugFunctionReturned(pass)
	return pass
}