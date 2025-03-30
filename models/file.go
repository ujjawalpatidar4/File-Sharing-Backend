package models

import (
	"time"
)

type File struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	UserID     uint      `json:"user_id"`
	Filename   string    `json:"filename"`
	Filepath   string    `json:"filepath"`
	Visibility string    `gorm:"default:private" json:"visibility"`
	SharedWith []string  `gorm:"type:text[]" json:"shared_with"`
	UploadedAt time.Time `json:"uploaded_at" gorm:"autoCreateTime"`
	ExpiresAt  time.Time `gorm:"index"`
}
