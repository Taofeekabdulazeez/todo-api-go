package model

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Title       string     `json:"title" gorm:"not null"`
	Description string     `json:"description"`
	Priority    string     `json:"priority" gorm:"default:'low'"`
	Completed   bool       `json:"completed" gorm:"default:false"`
	DueDate     *time.Time `json:"dueDate"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	UserID      uuid.UUID  `json:"userId" gorm:"type:uuid"`
}
