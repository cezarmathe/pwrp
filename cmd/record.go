package cmd

import (
	"github.com/cezarmathe/pwrp/recording"
	log "github.com/sirupsen/logrus"
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
	log.Trace("initializeRecorder(): ", "called")
	log.Debug("recorder: ", "creating recorder")
	recorder = recording.NewRecorder(config.Recording)
}

func runRecordCmd(cmd *cobra.Command, args []string) {
	log.Trace("runRecordCmd(): ", "called")

	recorder.Record()

	log.Trace("runRecordCmd(): ", "returned")
}
