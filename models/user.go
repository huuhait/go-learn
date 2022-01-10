package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        int64     `gorm:"primaryKey"`
	UID       string    `gorm:"type:character varying(10);not null"`
	Email     string    `gorm:"type:character varying(50);not null"`
	Password  string    `gorm:"type:text;not null"`
	Role      string    `gorm:"type:character varying(10);not null;default:member"`
	Status    string    `gorm:"type:character varying(10);not null;default:active"`
	CreatedAt time.Time `gorm:"type:timestamp(0)"`
	UpdatedAt time.Time `gorm:"type:timestamp(0)"`
	Musics    []Music
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.HashPassword()

	return nil
}

func (u *User) HashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)

	u.Password = string(bytes)
}

func (u *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// id, uid, email, password, role, state, created_at, updated_at
