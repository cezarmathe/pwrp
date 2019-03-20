/*
	PWRP - Personal Work Recorder Processor
	Copyright (C) 2019  Cezar Mathe <cezarmathe@gmail.com> [https://cezarmathe.com]

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Affero General Public License as published
	by the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU Affero General Public License for more details.

	You should have received a copy of the GNU Affero General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package cmd

import (
	cfg "github.com/cezarmathe/pwrp/config"
	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	configFileName string
	config         *cfg.Config
	configIsLoaded bool
)

func init() {
	config = new(cfg.Config)
	configIsLoaded = false
}

/*initConfig reads in config file and ENV variables if set.*/
func initConfig() {
	log.Trace("initConfig(): ", "called")

	/*Flag configurations*/
	log.Trace("initConfig(): ", "setting viper flag bindings")
	viper.BindPFlag("verbose", rootCmd.Flags().Lookup("verbose"))
	viper.SetEnvPrefix("PWRP")
	viper.BindEnv("debug")

	log.Trace("initConfig(): ", "reading environment variables")
	viper.AutomaticEnv() /*read in environment variables that match*/

	if viper.GetBool("verbose") {
		log.SetLevel(log.TraceLevel)
	} else if viper.GetBool("debug") {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	log.Debug("log level: ", log.GetLevel())

	/*Find home directory.*/
	log.Trace("initConfig(): ", "finding home directory")

	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	/*Value defaults*/
	log.Trace("initConfig(): ", "setting viper defaults")
	viper.SetDefault("verbose", false)
	viper.SetDefault("debug", false)

	if configFileName != "" {
		log.Debug("configuration file name: ", "using path from flag")
		/*Use config file from the flag.*/
		viper.SetConfigFile(configFileName)
	} else {
		log.Debug("configuration file name: ", "using default path")
		/*Search config in config directory with name "pwrp.toml".*/
		viper.SetConfigFile(home + "/.config/pwrp.toml")
	}

	/*If a config file is found, read it in.*/
	if err := viper.ReadInConfig(); err == nil {
		log.Debug("reading config file: ", "using "+viper.ConfigFileUsed())
	} else {
		log.Warn("reading config file:", "failed reading "+viper.ConfigFileUsed())
	}

	if err := viper.Unmarshal(config); err != nil {
		log.Warn("decoding config file: ", "failed to decode - ", err)
	}

	config.StoragePath = home + "/.local/share/pwrp"
	configIsLoaded = true

	log.Trace("initConfig(): ", "finished")
}
