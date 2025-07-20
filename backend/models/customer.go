package models

import (
	"time"

	"gorm.io/gorm"
)

// Customer represents a customer in the system
type Customer struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null"`
	TelNo     string         `json:"tel_no" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	// Orders []Order `json:"orders" gorm:"foreignKey:CustomerID"`
}
