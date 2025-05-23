package entities

import (
	"time"

	"github.com/jaayroots/todo-api/enums"
	"gorm.io/gorm"
)

type TodoStatus int

const (
	StatusNew TodoStatus = iota + 1
	StatusPending
	StatusInProgress
	StatusCompleted
)

type Todo struct {
	ID          uint             `gorm:"primaryKey"`
	Title       string           `gorm:"type:varchar(255);not null"`
	Description string           `gorm:"type:text"`
	Status      enums.TodoStatus `gorm:"type:int;default:1;not null"`
	DueDate     *time.Time

	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	CreatedBy uint  `gorm:"not null"`
	UpdatedBy uint  `gorm:"not null"`
	DeletedBy *uint `gorm:"column:deleted_by"`
}

func (t *Todo) GetCreatedBy() uint {
	return t.CreatedBy
}

func (t *Todo) GetUpdatedBy() uint {
	return t.UpdatedBy
}

func (t *Todo) GetDeletedBy() uint {
	if t.DeletedBy != nil {
		return *t.DeletedBy
	}
	return 0
}

func (t *Todo) BeforeCreate(tx *gorm.DB) error {
	if err := setBlameableFieldsBeforeCreate(tx, &t.CreatedBy, &t.UpdatedBy); err != nil {
		return err
	}
	return nil
}

func (t *Todo) BeforeUpdate(tx *gorm.DB) error {
	if err := setBlameableFieldsBeforeUpdate(tx, &t.UpdatedBy); err != nil {
		return err
	}
	return nil
}

func (t *Todo) BeforeDelete(tx *gorm.DB) error {
	return setBlameableFieldsBeforeDelete(tx, t, t.ID)
}