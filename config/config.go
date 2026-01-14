package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

// Postgres
type DBConfig struct {
	Name string `env:"POSTGRES_DB,required"`
	User string `env:"POSTGRES_USER,required"`
	Pass string `env:"POSTGRES_PASSWORD,required"`
	Host string `env:"POSTGRES_HOST" envDefault:"localhost"`
	SSL  string `env:"POSTGRES_SSL" envDefault:"disable"`
	Port string `env:"POSTGRES_PORT" envDefault:"5433"`
}

type TGConfig struct {
	TOKEN string `env:"TG_TOKEN,required"` // обязано существовать - Parse вернет ошибку
}

type RedisConfig struct {
	Addr     string `env:"REDIS_ADDR" envDefault:"localhost:6380"`
	Password string `env:"REDIS_PASSWORD,required"`
	DB       int    `env:"REDIS_DB" envDefault:"0"`
}

type MinioConfig struct {
	Endpoint 	string 	`env:"MINIO_ENDPOINT,required"`
	AccessKey 	string 	`env:"MINIO_ROOT_USER,required"`
	SecretKey 	string 	`env:"MINIO_ROOT_PASSWORD,required"`
	Secure 		bool	`env:"MINIO_SECURE" envDefault:"false"`
}

type Config struct {
	Db 	DBConfig
	App TGConfig
	Re 	RedisConfig
	Mi	MinioConfig
}

func Load() (*Config) {
	err := godotenv.Load() 
	if err != nil {
		log.Print("error - godotenv failed up env_file")
	}

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