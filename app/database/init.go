package database

import (
	"fmt"
	"strconv"
	"time"

	"github.com/aisalamdag23/MoneyMeExam/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Initialize - initializes database connection
func Initialize(cfg *config.Config) (*gorm.DB, error) {

	dbHost := cfg.DB.Host
	dbPort := cfg.DB.Port
	dbUser := cfg.DB.User
	dbPassword := cfg.DB.Pass
	dbName := cfg.DB.DBName
	dbConn := cfg.DB.Conns

	numberOfConnections, err := strconv.ParseInt(dbConn, 10, 64)

	if err != nil {
		numberOfConnections = 20 //default
	}

	dbURL := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable connect_timeout=5", dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(int(numberOfConnections))
	sqlDB.SetMaxIdleConns(int(numberOfConnections))
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	return db, err
}
