package main

import (
	"errors"
	"log"

	"github.com/Kosk0l/TgBotConverter/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg := config.Load()

	dsn := config.LoadDsn(cfg)

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