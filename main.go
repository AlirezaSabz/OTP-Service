package main

import (
	"log"
	"net/http"

	"github.com/AlirezaSabz/OTP-Service/internal/auth"
	"github.com/AlirezaSabz/OTP-Service/internal/cache"
	"github.com/AlirezaSabz/OTP-Service/internal/config"
	"github.com/AlirezaSabz/OTP-Service/internal/db"
	"github.com/AlirezaSabz/OTP-Service/internal/handlers"
	"github.com/AlirezaSabz/OTP-Service/internal/otp"
	"github.com/AlirezaSabz/OTP-Service/internal/user"
)

func main() {
	cfg := config.Load()

	postgres := db.Connect(cfg.PostgresDSN)
	redisCli := cache.Connect(cfg.RedisAddr)

	Repo := &user.Repository{DB: postgres}
	Repo.Migrate()
	userService := &user.Service{Repo: Repo}

	otpService := &otp.Service{Redis: redisCli}
	jwtService := &auth.JWTService{Secret: []byte(cfg.JWTSecret)}

	router := handlers.NewRouter(userService, otpService, jwtService)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
