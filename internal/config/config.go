package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type Config struct {
	DBUri          string `envconfig:"DB_URI" default:"postgres://homestead:homestead@localhost:54321/homestead"`
	SQLiteUri      string `envconfig:"SQLLITE_URI" default:"./ddd"`
	JWTSecret      string `envconfig:"JWT_SECRET" default:"FDERF$GRHHJ%TWETEHHYEH"`
	GrpcServerPort string `envconfig:"GRPC_SERVER_PORT" default:":8082"`
	MasterPassword string `envconfig:"MASTER_PASSWORD"`
	SslCertPath    string `envconfig:"SSL_CERT_PATH"`
	SslKeyPath     string `envconfig:"SSL_KEY_PATH"`
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Err(err).Msg("error load config from .env file")
	}

	config := Config{}
	err = envconfig.Process("", &config)
	if err != nil {
		return nil, errors.New("can not load config")
	}

	return &config, nil
}
