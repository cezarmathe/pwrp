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
	log.Debug("initializing")
	cobra.OnInitialize(initConfig)

	/*Persistent flags*/
	rootCmd.PersistentFlags().StringVarP(&configFileName, "config", "c", "", "config file (default is $HOME/.config/pwrp.toml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose level logging")
}

/*Execute adds all child commands to the root command and sets flags appropriately.
This is called by main.main(). It only needs to happen once to the rootCmd.*/
func Execute() {
	rootCmd.AddCommand(validateConfigCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runRootCmd(cmd *cobra.Command, args []string) {
	log.Debug("begin main command")

}
