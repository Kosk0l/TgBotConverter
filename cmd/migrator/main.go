package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/Kosk0l/TgBotConverter/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg := config.Load()

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Db.User,
		cfg.Db.Pass,
		cfg.Db.Host,
		cfg.Db.Port,
		cfg.Db.Name,
	)

	m, err := migrate.New("file://migrations", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("no migrations to apply — database already up to date")
			return
		}
		log.Fatal(err)
	}
}