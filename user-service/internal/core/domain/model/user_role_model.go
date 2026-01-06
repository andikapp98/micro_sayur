package model

import "time"

type UserRole struct {
	ID        int `gorm:"primaryKey"`
	UserID    int `gorm:"index"`
	RoleID    int `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (m UserRole) TableName() string {
    return "user_role"
}