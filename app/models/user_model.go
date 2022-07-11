package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

// User struct to describe User object.
type User struct {
	ID         uint32    `db:"id" json:"id" validate:"required"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	Name       string    `gorm:"type:varchar(255)" db:"name" json:"name" validate:"required,min=3,max=255"`
	Email      string    `gorm:"type:varchar(255);uniqueIndex" db:"email" json:"email" validate:"required,email,lte=255"`
	Password   string    `gorm:"type:varchar(255)" db:"password" json:"-" validate:"required,lte=255"`
	Phone      string    `gorm:"type:varchar(15);default:null" db:"phone" json:"phone"`
	PhotoID    uint32    `gorm:"TYPE:bigint REFERENCES files;default:null" db:"photo_id" json:"photo_id"`
	Photo      *File     `json:"photo"`
	UserStatus int       `gorm:"default:1" db:"user_status" json:"user_status" validate:"required,len=1"`
	UserRoleID uint32    `gorm:"default:3" db:"role_id" json:"role_id" validate:"required,len=1"`
	UserRole   UserRole  `json:"role"`
}

type UserRole struct {
	ID        uint32    `db:"id" json:"id" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Name      string    `gorm:"type:varchar(31)" db:"name" json:"name" validate:"required,lte=31"`
}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}
