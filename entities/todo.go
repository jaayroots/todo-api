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

	CreatedBy uint `gorm:"not null"`
	UpdatedBy uint `gorm:"not null"`
	DeletedBy uint `gorm:"column:deleted_by"`
}

func (t *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	if userID, ok := tx.Statement.Context.Value("userID").(uint); ok {
		t.CreatedBy = userID
		t.UpdatedBy = userID
	}
	return
}

func (t *Todo) BeforeUpdate(tx *gorm.DB) (err error) {
	if userID, ok := tx.Statement.Context.Value("userID").(uint); ok {
		t.UpdatedBy = userID
	}
	return
}

func (t *Todo) BeforeDelete(tx *gorm.DB) (err error) {
	if userID, ok := tx.Statement.Context.Value("userID").(uint); ok {
		tx.Statement.SetColumn("DeletedBy", userID)
	}
	return nil
}
