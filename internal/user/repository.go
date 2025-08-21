package user

import (
	"database/sql"
	"log"
	"time"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) Migrate() {
	_, err := r.DB.Exec(`CREATE TABLE IF NOT EXISTS users (
		phone TEXT PRIMARY KEY,
		registration_at TIMESTAMP NOT NULL
	)`)
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}
}

func (r *Repository) FindByPhone(phone string) (*User, error) {
	row := r.DB.QueryRow("SELECT phone, registration_at FROM users WHERE phone=$1", phone)
	var u User
	err := row.Scan(&u.Phone, &u.RegistrationAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &u, err
}

func (r *Repository) Create(phone string) error {
	_, err := r.DB.Exec("INSERT INTO users(phone, registration_at) VALUES($1, $2)", phone, time.Now())
	return err
}

func (r *Repository) List(limit, offset int, search string) ([]User, error) {
	query := "SELECT phone, registration_at FROM users"
	args := []any{}
	if search != "" {
		query += " WHERE phone LIKE $1"
		args = append(args, "%"+search+"%")
	}
	query += " ORDER BY registration_at DESC LIMIT $2 OFFSET $3"
	args = append(args, limit, offset)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		rows.Scan(&u.Phone, &u.RegistrationAt)
		users = append(users, u)
	}
	return users, nil
}
