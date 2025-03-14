package entity

import "time"

type User struct {
	ID        int        `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Email     string     `json:"email,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}
