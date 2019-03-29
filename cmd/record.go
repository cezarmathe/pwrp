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

	"github.com/cezarmathe/pwrp/recording"
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
	log.DebugFunctionCalled()
	defer log.DebugFunctionReturned()

	recording.InitLogging(log.GetParams())
	recorder = recording.NewRecorder(config.Recording)
}

func runRecordCmd(cmd *cobra.Command, args []string) {
	log.DebugFunctionCalled(*cmd, args)
	defer log.DebugFunctionReturned()

	if pass := runConfigValidation(); pass == false {
		log.Fatal("configuration did not pass the validation process")
	}
	log.Trace("configuration passed the validation process")

	log.Trace("initializing the recorder")
	initializeRecorder()

	log.Info("starting the recording process")
	recorder.Record()
}
