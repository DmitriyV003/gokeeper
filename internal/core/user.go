package core

import "time"

type User struct {
	ID        int64      `json:"id"`
	Login     string     `json:"login"`
	Password  string     `json:"-"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at"`
}
