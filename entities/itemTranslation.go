package entities

import (
	"time"

	"gorm.io/gorm"
)

type ItemTranslation struct {
	ID          uint   `gorm:"primaryKey"`
	ItemID      uint   `gorm:"index;not null"`
	Lang        string `gorm:"size:5;not null"`
	Title       string `gorm:"size:255;not null"`
	Description string `gorm:"type:text"`

	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	CreatedBy uint  `gorm:"not null"`
	UpdatedBy uint  `gorm:"not null"`
	DeletedBy *uint `gorm:"column:deleted_by"`
}

func (t *ItemTranslation) GetCreatedBy() uint {
	return t.CreatedBy
}

func (t *ItemTranslation) GetUpdatedBy() uint {
	return t.UpdatedBy
}

func (t *ItemTranslation) GetDeletedBy() uint {
	if t.DeletedBy != nil {
		return *t.DeletedBy
	}
	return 0
}

func (t *ItemTranslation) BeforeCreate(tx *gorm.DB) error {
	if err := setBlameableFieldsBeforeCreate(tx, &t.CreatedBy, &t.UpdatedBy); err != nil {
		return err
	}
	return nil
}

func (t *ItemTranslation) BeforeUpdate(tx *gorm.DB) error {
	if err := setBlameableFieldsBeforeUpdate(tx, &t.UpdatedBy); err != nil {
		return err
	}
	return nil
}

func (t *ItemTranslation) BeforeDelete(tx *gorm.DB) error {
	return setBlameableFieldsBeforeDelete(tx, t, t.ID)
}
