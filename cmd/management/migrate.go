package main

import (
	"fmt"
	"github.com/red-life/zone/cmd"
	"github.com/red-life/zone/internal/management"
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
	err = db.AutoMigrate(management.Zone{})
	if err != nil {
		log.Fatalln("error to migrate 'zones' table:", err)
	}
	err = db.AutoMigrate(management.Record{})
	if err != nil {
		log.Fatalln("error to migrate 'records' table:", err)
	}
	log.Println("done")
}
