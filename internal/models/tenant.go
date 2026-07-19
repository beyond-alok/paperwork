package models

import "time"

type Tenant struct {
	ID        string    `json:"id"         db:"id"`
	Name      string    `json:"name"       db:"name"`
	Username  string    `json:"username"      db:"username"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
