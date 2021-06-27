package models

import "github.com/google/uuid"

type File struct {
	FileId      uuid.UUID `json:"file_id" db:"file_id"`
	FileOwnerId uuid.UUID `json:"file_owner_id" db:"file_owner_id"`
	Filename    string    `json:"filename" db:"filename"`
	Description string    `json:"description,omitempty" db:"description"`
	Extension   string    `json:"extension" db:"extension"`
	Size        float64   `json:"size" db:"size"`
	Tags        []string  `json:"tags" db:"tags"`
	Base
}
