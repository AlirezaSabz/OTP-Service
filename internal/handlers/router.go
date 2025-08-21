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

	mux.HandleFunc("/request-otp", method(http.MethodPost, h.RequestOTP))
	mux.HandleFunc("/verify-otp", method(http.MethodPost, h.VerifyOTP))
	mux.HandleFunc("/user", method(http.MethodGet, h.GetUser))
	mux.HandleFunc("/users", method(http.MethodGet, h.ListUsers))

	return mux
}
func method(method string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		next(w, r)
	}
}
