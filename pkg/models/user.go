package models

import "time"

type User struct {
	ID        int64     `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Username  string    `json:"username,omitempty"`
	Password  string    `json:"password,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Type      Type      `json:"type,omitempty"`
}
