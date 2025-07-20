package models

import (
	"time"

	"gorm.io/gorm"
)

// Item represents a general item in the system
type Item struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null"`
	Type      string         `json:"type" gorm:"not null"` // 'pizza', 'beverage', 'other'
	UnitPrice float64        `json:"unit_price" gorm:"not null"`
	IsActive  bool           `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	// OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:ItemID"`
	// Beverages  []Beverage  `json:"beverages" gorm:"foreignKey:ItemID"`
	// Pizzas     []Pizza     `json:"pizzas" gorm:"foreignKey:ItemID"`
}
