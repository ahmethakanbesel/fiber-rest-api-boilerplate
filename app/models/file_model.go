package models

import (
	"time"
)

type File struct {
	ID        uint32    `db:"id" json:"id" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Path      string    `gorm:"type:varchar(255)" db:"path" json:"path" validate:"required,lte=255"`
	MimeType  string    `gorm:"type:varchar(31)" db:"mime_type" json:"mime_type"`
	Name      string    `gorm:"type:varchar(127)" db:"name" json:"name"`
	Size      uint32    `db:"size" json:"size"`
	OwnerID   uint32    `db:"owner_id" json:"owner_id" validate:"required"`
	Owner     User
	IsPublic  bool `gorm:"default:false" db:"is_public" json:"is_public"`
}
