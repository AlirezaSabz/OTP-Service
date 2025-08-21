package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresDSN string
	RedisAddr   string
	JWTSecret   []byte
}

func Load() Config {
	godotenv.Load()
	cfg := Config{
		PostgresDSN: os.Getenv("POSTGRES_DSN"),
		RedisAddr:   os.Getenv("REDIS_ADDR"),
		JWTSecret:   []byte(os.Getenv("JWT_SECRET")),
	}
	if cfg.PostgresDSN == "" || cfg.RedisAddr == "" || len(cfg.JWTSecret) == 0 {
		log.Fatal("missing environment variables")
	}
	return cfg
}
