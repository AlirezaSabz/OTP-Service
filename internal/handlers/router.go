package handlers

import (
	"net/http"

	"github.com/AlirezaSabz/OTP-Service/internal/auth"
	"github.com/AlirezaSabz/OTP-Service/internal/otp"
	"github.com/AlirezaSabz/OTP-Service/internal/user"
)

func NewRouter(userSvc *user.Service, otpSvc *otp.Service, jwtSvc *auth.JWTService) http.Handler {
	mux := http.NewServeMux()

	h := &Handler{
		UserService: userSvc,
		OTPService:  otpSvc,
		JWTService:  jwtSvc,
	}

	mux.HandleFunc("/request-otp", h.RequestOTP)
	mux.HandleFunc("/verify-otp", h.VerifyOTP)
	mux.HandleFunc("/user", h.GetUser)
	mux.HandleFunc("/users", h.ListUsers)

	return mux
}
