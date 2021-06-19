package models

import (
	"time"
)

type Base struct {
	//ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	//DeletedAt time.Time `json:"deleted_at" db:"deleted_at"`
}
