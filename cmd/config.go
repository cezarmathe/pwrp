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
	"github.com/cezarmathe/pwrp/gitops"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	cfg "github.com/cezarmathe/pwrp/config"
)

var (
	/*path to the configuration file*/
	configFilePath string

	/*the configuration object*/
	_config *cfg.Config

	/*configuration utility*/
	config *viper.Viper
)

func init() {
	/*initialize the configuration utility*/
	config = viper.New()

	_config = new(cfg.Config)
}

/*initConfig reads in the configuration file and ENV variables if set.*/
func initConfig() {
	log.DebugFunctionCalled()
	defer log.DebugFunctionReturned()

	/*flag configurations*/
	err := config.BindPFlag("verbose", rootCmd.Flags().Lookup("verbose"))
	if err != nil {
		log.FatalErr(err, "encountered an error when binding a flag")
	}

	/*env configurations*/
	config.SetEnvPrefix("pwrp")
	err = config.BindEnv("debug")
	err = config.BindEnv("verbose")
	if err != nil {
		log.FatalErr(err, "encountered an error when binding an environment variable")
	}

	/*read in environment variables that match*/
	config.AutomaticEnv()

	if config.GetBool("verbose") {
		log.SetLevel(logrus.TraceLevel)
		log.Trace("verbose logging level")
	} else {
		log.SetLevel(logrus.InfoLevel)
	}
	if config.GetBool("debug") {
		log.EnableDebug(true)
		log.Debug("debugging enabled")
	} else {
		log.EnableDebug(false)
	}

	/*find home directory.*/
	log.Trace("finding home directory")
	home, err := homedir.Dir()
	if err != nil {
		log.FatalErr(err, "encountered an error when trying to find the home directory")
	}

	/*set configuration defaults*/
	log.Trace("setting configuration defaults")
	viper.SetDefault("storage_path", home+"/.local/share/pwrp")
	viper.SetDefault("recording.repositories", []string{})
	viper.SetDefault("recording.protocol", gitops.DefaultProtocol)
	viper.SetDefault("recording.skips.missing_branch", false)
	viper.SetDefault("recording.skips.bad_url", false)
	viper.SetDefault("recording.skips.bad_protocol", false)
	viper.SetDefault("recording.skips.all", false)

	log.Trace("finding the configuration file")
	if configFilePath != "" {
		/*use the configuration file passed by the flag*/
		log.Trace("using the configuration file passed by flag")
	} else {
		/*search _config in _config directory with name "pwrp.toml".*/
		log.Trace("searching the configuration file in the default path")
		configFilePath = home + "/._config/pwrp.toml"

	}
	viper.SetConfigFile(configFilePath)

	/*if a configuration file is found, read it in.*/
	log.Trace("reading the configuration file")
	if err := config.ReadInConfig(); err == nil {
		log.Info("using " + config.ConfigFileUsed() + " as the configuration file")
	} else {
		log.FatalErr(err, "failed reading "+config.ConfigFileUsed())
	}

	/*load the configuration into the _config object*/
	log.Trace("loading the configuration into the _config object")
	if err := viper.Unmarshal(_config); err != nil {
		log.FatalErr(err, "failed to the decode the configuration file")
	}
}
