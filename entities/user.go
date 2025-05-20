package entities

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Email     string    `gorm:"type:varchar(128);unique;not null;"`
	Password  string    `gorm:"type:varchar(256);unique;not null;"`
	FirstName string    `gorm:"type:varchar(128);not null;"`
	LastName  string    `gorm:"type:varchar(128);not null;"`
	Avatar    string    `gorm:"type:varchar(256);default:'';"`
	IsDeleted bool      `gorm:"not null;default:false;"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime;"`

	CreatedBy []Todo `gorm:"foreignKey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UpdatedBy []Todo `gorm:"foreignKey:UpdatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DeletedBy []Todo `gorm:"foreignKey:DeletedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (u User) GetID() uint {
	return u.ID
}
