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
	"fmt"
	"os"

	"github.com/cezarmathe/pwrp/smartlogger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

/*The main command*/
var rootCmd = &cobra.Command{
	Use:   "pwrp",
	Short: "A utility for recording and processing personal work",
	Run:   runRootCmd,
}

var (
	log *smartlogger.SmartLogger
)

func init() {
	log = smartlogger.NewSmartLogger(false, logrus.InfoLevel)

	/*Run the cobra initialization process*/
	cobra.OnInitialize(initConfig)

	/*Persistent flags*/
	rootCmd.PersistentFlags().StringVarP(&configFilePath, "config", "c", "", "config file (default is $HOME/.config/pwrp.toml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose level logging")
}

/*Execute adds all child commands to the root command and sets flags appropriately.
This is called by main.main(). It only needs to happen once to the rootCmd.*/
func Execute() {
	log.Debug("called")

	log.Debug("adding additional commands to the root command")
	rootCmd.AddCommand(validateConfigCmd)
	rootCmd.AddCommand(recordCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log.Debug("returned")
}

func runRootCmd(cmd *cobra.Command, args []string) {
	log.Debug("called")

	// FIXME 29/03 cezarmathe: if the configuration failed to load, do not continue
	log.Trace("initializing the recorder")
	initializeRecorder()

	log.Info("starting the recording process")
	if success := recorder.Record(); success == false {
		log.Fatal("cannot continue due to recording failure")
	}
	log.Info("recording was successful")

	log.Debug("returned")
}
