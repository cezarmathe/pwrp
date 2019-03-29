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
	cfg "github.com/cezarmathe/pwrp/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	validateConfigCmd = &cobra.Command{
		Use:   "validate-config",
		Short: "Validate the configuration file",
		Run:   runValidateConfigCmd,
	}
)

func runValidateConfigCmd(cmd *cobra.Command, args []string) {
	logrus.Trace("runValidateConfig(): ", "called")
	cfg.ValidateConfig(config)
	logrus.Trace("runValidateConfig: ", "returned")
}
