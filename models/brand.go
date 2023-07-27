package models

import (
	"time"
)

type (
	// Brand
	Brand struct {
		ID        uint      `gorm:"primary_key" json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Phones    []Phone   `json:"-"`
		Models    []Model   `json:"-"`
	}
)
