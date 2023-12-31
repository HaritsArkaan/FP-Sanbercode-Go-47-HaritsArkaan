package models

import (
	"time"
)

type (
	// Review
	Review struct {
		ID        uint      `gorm:"primary_key" json:"id"`
		Text      string    `json:"text"`
		UserID    uint      `json:"user_id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		User      User      `json:"-"`
		Phone     Phone     `json:"-"`
	}
)
