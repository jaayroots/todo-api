package entities

import "time"

type Session struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement"`
	UserID    uint64    `gorm:"not null;"`
	Token     string    `gorm:"type:text;not null;"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
	ExpiresAt time.Time `gorm:"not null;autoCreateTime;"`
	IpAddress string    `gorm:"type:varchar(128);not null;"`
}
