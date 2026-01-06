package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

// Postgres
type DBConfig struct {
	Host string `env:"DB_HOST" envDefault:"localhost"`
	Port string `env:"DB_PORT" envDefault:"5433"`
	User string `env:"DB_USER,required"`
	Pass string `env:"DB_PASS,required"`
	Name string `env:"DB_NAME,required"`
	SSL  string `env:"DB_SSL" envDefault:"disable"`
}

type TGConfig struct {
	TOKEN string `env:"TG_TOKEN,required"` // обязано существовать - Parse вернет ошибку
}

type RedisConfig struct {
	Addr     string `env:"REDIS_ADDR" envDefault:"localhost:6380"`
	Password string `env:"REDIS_PASSWORD,required"`
	DB       int    `env:"REDIS_DB" envDefault:"0"`
}

type Config struct {
	Db 	DBConfig
	App TGConfig
	Re 	RedisConfig
}

func Load() (*Config) {
	_ = godotenv.Load()

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	return &cfg
}

//====================================================================================================

func LoadDsn(cfg *Config) (string) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Db.User,
		cfg.Db.Pass,
		cfg.Db.Host,
		cfg.Db.Port,
		cfg.Db.Name,
	)
	return dsn
}