package models

import (
	"time"
)

type (
	// Rating
	Rating struct {
		ID        uint      `gorm:"primary_key" json:"id"`
		Rating    uint      `json:"rating"`
		UserID    uint      `json:"user_id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		User      User      `json:"-"`
		Phone     Phone     `json:"-"`
	}
)
