package database

import (
	"fmt"
	"log"

	"github.com/aisalamdag23/MoneyMeExam/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(cfg *config.Config) error {

	dbHost := cfg.DB.Host
	dbPort := cfg.DB.Port
	dbUser := cfg.DB.User
	dbPassword := cfg.DB.Pass
	dbName := cfg.DB.DBName

	dbURL := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	m, err := migrate.New("file://db/migrations", dbURL)

	if err != nil {
		return err
	}

	if err = m.Up(); err != nil {
		log.Printf("db migrate up failed: %v", err.Error())
		if err.Error() != "no change" {
			return err
		}
	}

	return nil
}
