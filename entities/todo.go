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
	ID          int              `gorm:"primaryKey"`
	Title       string           `gorm:"type:varchar(255);not null"`
	Description string           `gorm:"type:text"`
	Status      enums.TodoStatus `gorm:"type:int;default:1;not null"`
	DueDate     *time.Time
	CreatedAt   time.Time      `gorm:"not null;autoCreateTime;"`
	UpdatedAt   time.Time      `gorm:"not null;autoUpdateTime;"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	CreatedBy   int            `gorm:"not null"`
	UpdatedBy   int            `gorm:"not null"`
	DeletedBy   int            `gorm:"default:null"`
}

func (t *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	if userID, ok := tx.Statement.Context.Value("userID").(int); ok {
		t.CreatedBy = userID
		t.UpdatedBy = userID
	}
	return
}

func (t *Todo) BeforeUpdate(tx *gorm.DB) (err error) {
	if userID, ok := tx.Statement.Context.Value("userID").(int); ok {
		t.UpdatedBy = userID
	}
	return
}

func (t *Todo) BeforeDelete(tx *gorm.DB) (err error) {
	if userID, ok := tx.Statement.Context.Value("userID").(int); ok {
		err := tx.Model(&Todo{}).
			Where("id = ?", t.ID).
			Updates(map[string]interface{}{
				"deleted_by": userID,
				"updated_at": time.Now(),
			}).Error
		if err != nil {
			return err
		}
	}
	return
}
