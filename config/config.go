package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	AppEnv string `required:"true" default:"prod" envconfig:"APP_ENV"`
	DB     struct {
		Username string `required:"true" envconfig:"MONGO_USERNAME"`
		Password string `required:"true" envconfig:"MONGO_PASSWORD"`
		DBName   string `required:"true" envconfig:"MONGO_DB_NAME"`
		Host     string `required:"true" envconfig:"MONGO_HOST"`
		Port     string `required:"true" envconfig:"MONGO_PORT"`
	}
}

func NewConfig() *Config {
	var cfg = new(Config)

	err := envconfig.Process("crew-app", cfg)
	if err != nil {
		log.Fatalf("could not process config: %v", err)
	}

	return cfg
}

func (c *Config) GetDBUrl() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/admin?authSource=admin",
		c.DB.Username,
		c.DB.Password,
		c.DB.Host,
		c.DB.Port,
	)
}
