package models

import (
	"time"
)

type (
	// Phone
	Phone struct {
		ID        uint      `gorm:"primary_key" json:"id"`
		BrandlID  uint      `json:"brand_id"`
		ColorID   uint      `json:"color_id"`
		ModelID   uint      `json:"model_id"`
		Price     string    `json:"price"`
		Storage   uint      `json:"storage"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Brand     Brand     `json:"-"`
		// Color     Color     `json:"-"`
		// Model Model `json:"-"`
	}
)
