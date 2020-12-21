package models

import "time"

type Model struct {
	ID        int       `gorose:"id"`
	CreatedAt time.Time `gorose:"created_at"`
	UpdatedAt time.Time `gorose:"updated_at"`
	// DeletedAt int `gorose:"deleted_at"`
}
