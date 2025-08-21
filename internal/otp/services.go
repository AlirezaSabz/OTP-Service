package otp

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/redis/go-redis/v9"
)

type Service struct {
	Redis *redis.Client
}

const (
	otpTTL       = 2 * time.Minute
	rateLimitTTL = 10 * time.Minute
	rateLimitMax = 3
)

func generateCode() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func (s *Service) Generate(ctx context.Context, phone string) error {
	key := fmt.Sprintf("otp_requests:%s", phone)
	count, _ := s.Redis.Get(ctx, key).Int()
	if count >= rateLimitMax {
		return fmt.Errorf("too many requests, try later")
	}
	s.Redis.Incr(ctx, key)
	s.Redis.Expire(ctx, key, rateLimitTTL)

	code := generateCode()
	otpKey := fmt.Sprintf("otp:%s", phone)
	if err := s.Redis.Set(ctx, otpKey, code, otpTTL).Err(); err != nil {
		return err
	}

	log.Printf("OTP for %s is %s\n", phone, code)
	return nil
}

func (s *Service) Verify(ctx context.Context, phone, code string) (bool, error) {
	otpKey := fmt.Sprintf("otp:%s", phone)
	stored, err := s.Redis.Get(ctx, otpKey).Result()
	if err == redis.Nil {
		return false, fmt.Errorf("otp expired or not found")
	} else if err != nil {
		return false, err
	}

	if stored != code {
		return false, fmt.Errorf("invalid otp")
	}

	s.Redis.Del(ctx, otpKey)
	return true, nil
}
