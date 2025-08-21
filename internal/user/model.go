package user

import "time"

type User struct {
	Phone          string    `json:"phone"`
	RegistrationAt time.Time `json:"registration_at"`
}
