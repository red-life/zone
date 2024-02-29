package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/red-life/zone/cmd"
	"github.com/red-life/zone/internal/management"
	"github.com/red-life/zone/internal/management/adapters/http"
	postgresRepo "github.com/red-life/zone/internal/management/adapters/postgres"
	"github.com/red-life/zone/pkg/mq"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

func newGormDB(postgresCfg cmd.Postgres) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		postgresCfg.Host, postgresCfg.Username, postgresCfg.Password,
		postgresCfg.DB, postgresCfg.Port)
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
	return db
}

func newRedisClient(redisCfg cmd.Redis) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		log.Fatalf("failed to ping redis: %s", err)
	}
	return rdb
}

func main() {
	cfg := cmd.ParseManagementConfig()
	db := newGormDB(cfg.Postgres)
	repo := postgresRepo.NewManagementRepository(db)
	rdb := newRedisClient(cfg.Redis)
	redisMQ := mq.NewRedisMessageQueue(rdb)
	service := management.NewManagementService(repo, redisMQ)
	api := http.NewAPI(service)
	engine := gin.Default()
	authMiddleware := http.BasicAuth(cfg.Username, cfg.Password)
	httpServer := http.NewHTTPServer(api, engine, cfg.Listen, authMiddleware)
	httpServer.RegisterRoutes()
	log.Fatalln(httpServer.Run())
}
