package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AlirezaSabz/OTP-Service/internal/auth"
	"github.com/AlirezaSabz/OTP-Service/internal/otp"
	"github.com/AlirezaSabz/OTP-Service/internal/user"
)

type Handler struct {
	UserService *user.Service
	OTPService  *otp.Service
	JWTService  *auth.JWTService
}

func (h *Handler) RequestOTP(w http.ResponseWriter, r *http.Request) {
	phone := r.URL.Query().Get("phone")
	if phone == "" {
		http.Error(w, "phone required", http.StatusBadRequest)
		return
	}
	//TODO : we can add some constraint for Phone Number

	if err := h.OTPService.Generate(context.Background(), phone); err != nil {
		http.Error(w, err.Error(), http.StatusTooManyRequests)
		return
	}
	w.Write([]byte("OTP generated (check server logs)"))
}

func (h *Handler) VerifyOTP(w http.ResponseWriter, r *http.Request) {
	phone := r.URL.Query().Get("phone")
	code := r.URL.Query().Get("otp")
	if phone == "" || code == "" {
		http.Error(w, "phone and otp required", http.StatusBadRequest)
		return
	}

	ok, err := h.OTPService.Verify(context.Background(), phone, code)
	if err != nil || !ok {
		http.Error(w, "invalid or expired otp", http.StatusUnauthorized)
		return
	}

	if err := h.UserService.RegisterIfNotExists(phone); err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}

	token, _ := h.JWTService.CreateToken(phone)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	phone := r.URL.Query().Get("phone")
	if phone == "" {
		http.Error(w, "phone required", http.StatusBadRequest)
		return
	}

	u, err := h.UserService.Get(phone)
	if err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}
	if u == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	search := r.URL.Query().Get("search")

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	users, err := h.UserService.List(limit, offset, search)
	if err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
