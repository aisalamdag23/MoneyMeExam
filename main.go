package main

import (
	"errors"
	"log"

	"github.com/aisalamdag23/MoneyMeExam/api/handlers"
	"github.com/aisalamdag23/MoneyMeExam/api/middleware"
	"github.com/aisalamdag23/MoneyMeExam/app/database"
	"github.com/aisalamdag23/MoneyMeExam/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// load configs
	cfg, err := config.Load()
	if err != nil || cfg == nil {
		if err == nil {
			err = errors.New("config is nil")
		}
		log.Fatalf("unable to load configurations: %v", err)
		return
	}

	// db migration
	err = database.Migrate(cfg)
	if err != nil {
		log.Fatalf("unable to migrate db: %v", err)
		return
	}

	// init db connection
	db, err := database.Initialize(cfg)
	if err != nil {
		log.Fatalf("unable to initialize db: %v", err)
		return
	}

	// api init
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(middleware.ContextMiddleware())
	router.Use(middleware.DatabaseMiddleware(db))

	router.POST("/request-loan", handlers.PostLoanRequest)

	router.Run("localhost:8081")
}
