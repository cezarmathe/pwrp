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
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	cfg "github.com/cezarmathe/pwrp/config"
)

var (
	/*path to the configuration file*/
	configFilePath string

	/*the configuration object*/
	config *cfg.Config
)

func init() {
	config = new(cfg.Config)
}

/*initConfig reads in config file and ENV variables if set.*/
func initConfig() {
	log.DebugFunctionCalled()
	defer log.DebugFunctionReturned()

	/*flag configurations*/
	err := viper.BindPFlag("verbose", rootCmd.Flags().Lookup("verbose"))
	if err != nil {
		log.FatalErr(err, "encountered an error when binding a flag")
	}

	/*env configurations*/
	viper.SetEnvPrefix("PWRP")
	err = viper.BindEnv("debug")
	if err != nil {
		log.FatalErr(err, "encountered an error when binding an environment variable")
	}

	/*read in environment variables that match*/
	viper.AutomaticEnv()

	if viper.GetBool("verbose") {
		log.SetLevel(logrus.TraceLevel)
		log.Trace("verbose logging level")
	} else {
		log.SetLevel(logrus.InfoLevel)
	}
	if viper.GetBool("debug") {
		log.EnableDebug(true)
		log.Debug("debugging enabled")
	}

	/*find home directory.*/
	log.Trace("finding home directory")

	home, err := homedir.Dir()
	if err != nil {
		log.FatalErr(err, "encountered an error when trying to find the home directory")
	}

	log.Trace("finding the configuration file")
	if configFilePath != "" {
		/*use the configuration file passed by the flag*/
		log.Trace("using the configuration file passed by flag")
		viper.SetConfigFile(configFilePath)
	} else {
		/*search config in config directory with name "pwrp.toml".*/
		log.Trace("searching the configuration file in the default path")
		viper.SetConfigFile(home + "/.config/pwrp.toml")
	}

	/*if a config file is found, read it in.*/
	log.Trace("reading the configuration file")
	if err := viper.ReadInConfig(); err == nil {
		log.Info("using " + viper.ConfigFileUsed() + " as the configuration file")
	} else {
		log.FatalErr(err, "failed reading "+viper.ConfigFileUsed())
	}

	/*load the configuration into the config object*/
	log.Trace("loading the configuration into the config object")
	if err := viper.Unmarshal(config); err != nil {
		log.FatalErr(err, "failed to the decode the configuration file")
	}

	/*setting configuration defaults*/
	config.StoragePath = home + "/.local/share/pwrp"
}
