package cmd

import (
	cfg "github.com/cezarmathe/pwrp/config"
	log "github.com/sirupsen/logrus"
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
	log.Trace("runValidateConfig(): ", "called")
	cfg.ValidateConfig(config)
	log.Trace("runValidateConfig: ", "returned")
}
