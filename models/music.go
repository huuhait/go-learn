package models

import "time"

type Music struct {
	ID        int64     `gorm:"primaryKey"`
	UserID    int64     `gorm:"type:bigint"`
	CreatedAt time.Time `gorm:"type:timestamp(0)"`
	UpdatedAt time.Time `gorm:"type:timestamp(0)"`
	User      User
}
