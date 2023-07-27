package models

import (
	"time"
)

type (
	// Model
	Model struct {
		ID        uint      `gorm:"primary_key" json:"id"`
		Name      string    `json:"name"`
		BrandlID  uint      `json:"brand_id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Brand     Brand     `json:"-"`
		// Phones    []Phone   `json:"-"`
	}
)
