package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Taskmark struct {
	gorm.Model
	Id        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	TaskName  string    `json:"task_name"`
	TotalDays int       `json:"total_days"`
	FromDate  string    `json:"from_date" gorm:"type:date"`
	TaskDays  string    `json:"task_days" gorm:"type:text"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}
