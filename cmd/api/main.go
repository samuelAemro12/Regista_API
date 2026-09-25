package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"

	"github.com/regista/regista-api/internal/config"
	"github.com/regista/regista-api/internal/db"
	"github.com/regista/regista-api/internal/middleware"
	"github.com/regista/regista-api/internal/routes"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load configuration: %v", err)
	}

	ctx := context.Background()

	var postgresPool *pgxpool.Pool
	pool, err := db.NewPostgresPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Printf("WARNING: PostgreSQL connection failed: %v (continuing without database)", err)
	} else {
		postgresPool = pool
		defer postgresPool.Close()
	}

	var redisClient *redis.Client
	client, err := db.NewRedisClient(ctx, cfg.RedisURL)
	if err != nil {
		log.Printf("WARNING: Redis connection failed: %v (continuing without Redis)", err)
	} else {
		redisClient = client
		defer redisClient.Close()
	}

	router := routes.NewRouter()
	router.Use(middleware.Logger(), middleware.CORS())
	routes.RegisterAPI(router)

	log.Printf("starting API server on :%s (%s)", cfg.Port, cfg.AppEnv)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("run API server: %v", err)
	}
}
