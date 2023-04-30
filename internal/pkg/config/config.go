package config

import (
	"context"
	"log"

	"github.com/kelseyhightower/envconfig"
)

type ServiceConfig struct {
	APIPort int `envconfig:"PORT" default:"8181"`
}

type DBConfig struct {
	DB_URL      string `envconfig:"DB_URL" default:"localhost:3306"`
	DB_USERNAME string `envconfig:"DB_USERNAME" required:"true"`
	DB_PASSWORD string `envconfig:"DB_PASSWORD" required:"true"`
}

// GetServiceConfig method to fetch the ServiceConfig
func GetServiceConfig(ctx context.Context) ServiceConfig {
	log.Println("Fetching Service configs")
	config := ServiceConfig{}

	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatalln(ctx, "Failed fetching service configs")
		panic(err)
	}
	return config
}

// GetDBConfig get db env vars or error
func GetDBConfig(ctx context.Context) DBConfig {
	dbConfig := DBConfig{}
	err := envconfig.Process("", &dbConfig)
	if err != nil {
		log.Fatalln(ctx, "Failed fetching db configs")
		panic(err)
	}
	return dbConfig
}
