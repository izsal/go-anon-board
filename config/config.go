package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	config *viper.Viper
}

func NewConfig() *Config {
	c := new(Config)
	c.config = readConfig()
	return c
}

func (c *Config) Get() *viper.Viper {
	if c.config == nil {
		log.Fatal("config is intialized")
	}
	return c.config
}

func readConfig() *viper.Viper {
	log.Info("reading environtment variables")
	v := viper.New()
	v.AutomaticEnv()
	env := v.GetString("ENVIRONTMENT")
	if env == "" {
		env = "local"
	}
	log.Infof("ENVIRONMENT: %s", env)
	v.SetConfigName(env)
	v.SetConfigType("yaml")
	v.AddConfigPath("config")

	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("error reading config file or env variabale : '%s'", err.Error())
	}

	return v
}
