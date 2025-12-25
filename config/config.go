package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host string `env:"DB_HOST" envDefault:"localhost"`
	Port string `env:"DB_PORT" envDefault:"5432"`
	User string `env:"DB_USER,required"`
	Pass string `env:"DB_PASS,required"`
	Name string `env:"DB_NAME,required"`
	SSL  string `env:"DB_SSL" envDefault:"disable"`
}

type TGConfig struct {
	TOKEN string `env:"TG_TOKEN,required"` // обязано существовать - Parse вернет ошибку
}

type Config struct {
	Db 	DBConfig
	App TGConfig
}

func Load() (*Config) {
	_ = godotenv.Load()

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	return &cfg
}