package models

import "time"

type Order struct {
	Id          int64     `json:"id,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	Hardware    string    `json:"hardware,omitempty"`
	TypeFailure string    `json:"type_failure,omitempty"`
	Description string    `json:"description,omitempty"`
	ClientId    int64     `json:"client,omitempty"`
	ExecutorId  int64     `json:"executor,omitempty"`
	Status      Status    `json:"status,omitempty"`
}
