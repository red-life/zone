package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/red-life/zone/cmd"
	"github.com/red-life/zone/internal/management"
	"github.com/red-life/zone/internal/management/adapters/http"
	postgresRepo "github.com/red-life/zone/internal/management/adapters/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	cfg := cmd.ParseManagementConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Postgres.Host, cfg.Postgres.Username, cfg.Postgres.Password,
		cfg.Postgres.DB, cfg.Postgres.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}
	sqlDB, _ := db.DB()
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("failed to ping db: %s\n", err)
	}
	repo := postgresRepo.NewManagementRepository(db)
	service := management.NewManagementService(repo)
	api := http.NewAPI(service)
	engine := gin.Default()
	authMiddleware := http.BasicAuth(cfg.Username, cfg.Password)
	httpServer := http.NewHTTPServer(api, engine, cfg.Listen, authMiddleware)
	httpServer.RegisterRoutes()
	log.Fatalln(httpServer.Run())
}
