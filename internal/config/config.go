package config

import (
	"github.com/rs/zerolog/log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {

	BigtableEmulatorHost      string        `mapstructure:"BIGTABLE_EMULATOR_HOST"`
}

func Load() (cfg *Config, err error) {
	// load the external ENV
	viper.AutomaticEnv() // read in environment variables that match
	// load default config file if existed
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	viper.SetConfigType("env")
	viper.AddConfigPath(dir)
	err = viper.ReadInConfig()
	if err != nil {
		log.Debug().Msg(err.Error())
	}
	err = viper.Unmarshal(&cfg)
	return cfg, err
}
