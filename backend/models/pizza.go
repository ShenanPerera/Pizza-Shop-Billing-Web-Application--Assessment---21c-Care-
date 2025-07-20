package models

import (
	"time"

	"gorm.io/gorm"
)

// Pizza represents a pizza item
type Pizza struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	ItemID    uint           `json:"item_id" gorm:"not null"`
	Name      string         `json:"name" gorm:"not null"`
	Size      string         `json:"size" gorm:"not null"`
	BaseType  string         `json:"base_type" gorm:"not null"`
	Price     float64        `json:"price" gorm:"not null"`
	IsActive  bool           `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	// Item     Item             `json:"item" gorm:"foreignKey:ItemID"`
	// Toppings []PizzaToppings  `json:"toppings" gorm:"foreignKey:PizzaID"`
}
