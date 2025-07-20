package models

import (
	"time"

	"gorm.io/gorm"
)

// Beverage represents a beverage item (modified based on ER diagram)
type Beverage struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	ItemID     uint           `json:"item_id" gorm:"not null"`
	BeverageID uint           `json:"beverage_id" gorm:"not null"` // References beverages lookup table
	Name       string         `json:"name" gorm:"not null"`
	Size       string         `json:"size" gorm:"not null"`
	Price      float64        `json:"price" gorm:"not null"`
	IsActive   bool           `json:"is_active" gorm:"default:true"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	// Item Item `json:"item" gorm:"foreignKey:ItemID"`
	// BeverageDetail BeverageDetail `json:"beverage_detail" gorm:"foreignKey:BeverageID"`
}
