package entities

import (
	"time"
)

type User struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement"`
	Email     string    `gorm:"type:varchar(128);unique;not null;"`
	Password  string    `gorm:"type:varchar(256);unique;not null;"`
	FirstName string    `gorm:"type:varchar(128);not null;"`
	LastName  string    `gorm:"type:varchar(128);not null;"`
	Avatar    string    `gorm:"type:varchar(256);default:'';"`
	IsDeleted  bool      `gorm:"not null;default:false;"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime;"`
}
