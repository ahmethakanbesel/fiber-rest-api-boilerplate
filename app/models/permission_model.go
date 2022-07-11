package models

import (
	"time"
)

type Permission struct {
	ID           uint32    `db:"id" json:"id" validate:"required"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	Name         string    `gorm:"type:varchar(127)" db:"name" json:"name" validate:"required,lte=127"`
	DisplayName  string    `gorm:"type:varchar(127)" db:"display_name" json:"display_name" validate:"lte=127"`
	Description  string    `gorm:"type:varchar(255)" db:"description" json:"description" validate:"lte=255"`
	Group        string    `gorm:"type:varchar(127)" db:"group" json:"group" validate:"required,lte=127"`
	Restrictions string    `gorm:"type:varchar(255)" db:"restrictions" json:"restrictions" validate:"lte=255"`
}

type RolePermission struct {
	ID           uint32    `db:"id" json:"id" validate:"required"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	UserRoleID   uint32    `db:"user_role" json:"user_role" validate:"required"`
	UserRole     UserRole
	PermissionID uint32 `db:"permission_id" json:"permission_id" validate:"required"`
	Permission   Permission
	Name         string `gorm:"type:varchar(127)" db:"name" json:"name" validate:"lte=127"`
}
