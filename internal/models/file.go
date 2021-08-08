package models

import "github.com/google/uuid"

type File struct {
	FileId      uuid.UUID `json:"file_id" db:"file_id"` // ?
	FileOwnerId uuid.UUID `json:"file_owner_id" db:"file_owner_id"`
	File        *UploadInput
	Description string   `json:"description,omitempty" db:"description"`
	Tags        []string `json:"tags" db:"tags"`
	Base
}
