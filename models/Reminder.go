package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Reminder struct {
	gorm.Model
	Id        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserID    uuid.UUID `json:"user_id"`
	Remind    string    `json:"remind"`
	Time      time.Time `json:"time"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}
