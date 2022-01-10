package entities

import "time"

type User struct {
	ID        int64     `json:"id"`
	UID       string    `json:"uid"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
