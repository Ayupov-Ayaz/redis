package mysql

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Host     string `envconfig:"MYSQL_HOST" validate:"required"`
	Port     int    `envconfig:"MYSQl_PORT" validate:"required"`
	Username string `envconfig:"MYSQl_USER" validate:"required"`
	Password string `envconfig:"MYSQL_PASSWORD" validate:"required"`
	Database string `envconfig:"MYSQL_DATABASE" validate:"required"`
}

func (cfg *config) GetDriver() string {
	return "mysql"
}

func (cfg *config) GetDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
}

func NewConfig() (*config, error) {
	if err := godotenv.Load("./components/mysql/.env"); err != nil {
		return nil, err
	}

	c := new(config)
	if err := envconfig.Process("MYSQL", c); err != nil {
		return nil, err
	}
	return c, nil
}