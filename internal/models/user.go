package models

import "time"

type Role string

type User struct {
	ID           string    `json:"id"         db:"id"`
	Name         string    `json:"name"       db:"name"`
	Email        string    `json:"email"      db:"email"`
	PasswordHash string    `json:"-"          db:"password_hash"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}


