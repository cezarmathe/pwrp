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
