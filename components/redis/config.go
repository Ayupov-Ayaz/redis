package redis

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config interface {
	GetAddr() string
}

type config struct {
	Host string `envconfig:"REDIS_HOST" validate:"required"`
	Port uint   `envconfig:"REDIS_PORT" validate:"required"`
}

func (cfg *config) GetAddr() string {
	return fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
}

func NewConfig() (Config, error) {
	if err := godotenv.Load("./components/redis/.env"); err != nil {
		return nil, err
	}

	c := new(config)
	if err := envconfig.Process("REDIS", c); err != nil {
		return nil, err
	}
	return c, nil
}
