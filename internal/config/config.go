package config

import (
	"flag"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DBUri string `json:"db_uri"`
}

func Load() (*Config, error) {
	path, err := getFilePath()
	if err != nil {
		return nil, errors.Wrap(err, "can not get config path")
	}

	if path == "" {
		return nil, errors.New("can not load config from empty path")
	}

	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("json")

	err = v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := Config{}
	err = v.Unmarshal(&config)
	if err != nil {
		return nil, errors.Wrap(err, "can not unmarshal config from file to struct")
	}
	log.Println(config)

	return &config, nil
}

func getFilePath() (string, error) {
	flag.String("config", "./config/config.json", "path to config file")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		return "", err
	}
	viper.AutomaticEnv()

	return viper.GetString("config"), nil
}
