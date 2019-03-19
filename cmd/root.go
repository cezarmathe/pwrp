package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

/*The main command*/
var rootCmd = &cobra.Command{
	Use:   "pwrp",
	Short: "A utility for recording and processing personal work",
	Run:   runRootCmd,
}

func init() {
	log.Trace("cmd.init(): ", "called")
	log.Debug("cmd init: ", "initializing")
	cobra.OnInitialize(initConfig)

	/*Persistent flags*/
	rootCmd.PersistentFlags().StringVarP(&configFileName, "config", "c", "", "config file (default is $HOME/.config/pwrp.toml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose level logging")

	log.Trace("cmd.init(): ", "returned")
}

/*Execute adds all child commands to the root command and sets flags appropriately.
This is called by main.main(). It only needs to happen once to the rootCmd.*/
func Execute() {
	log.Trace("cmd.Execute(): ", "called")

	log.Debug("execute root command: ", "adding commands")
	rootCmd.AddCommand(validateConfigCmd)
	rootCmd.AddCommand(recordCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log.Trace("cmd.Execute(): ", "returned")
}

func runRootCmd(cmd *cobra.Command, args []string) {
	log.Trace("runRootCmd(): ", "called")

	log.Debug("root cmd: ", "running recorder.Record()")
	if success := recorder.Record(); success == false {
		log.Fatal("cannot continue")
	}
	log.Info("can continue")

	log.Trace("runRootCmd(): ", "returned")
}
