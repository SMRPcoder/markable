package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Taskmark struct {
	gorm.Model
	Id        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	TotalDays string    `json:"total_days"`
	TaskDays  string    `json:"task_days"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}
