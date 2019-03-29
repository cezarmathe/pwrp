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
	"github.com/cezarmathe/pwrp/recording"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	recorder  *recording.Recorder
	recordCmd = &cobra.Command{
		Use:   "record",
		Short: "Run the recording process alone",
		Run:   runRecordCmd,
	}
)

func initializeRecorder() {
	logrus.Trace("initializeRecorder(): ", "called")
	logrus.Debug("recorder: ", "creating recorder")
	recorder = recording.NewRecorder(config.Recording, logrus.New())
}

func runRecordCmd(cmd *cobra.Command, args []string) {
	logrus.Trace("runRecordCmd(): ", "called")

	initializeRecorder()
	recorder.Record()

	logrus.Trace("runRecordCmd(): ", "returned")
}
