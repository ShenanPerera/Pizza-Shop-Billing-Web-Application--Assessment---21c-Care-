package models

import (
	"time"

	"gorm.io/gorm"
)

// Topping represents available toppings
type Topping struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	ToppingID uint           `json:"topping_id" gorm:"not null"`
	Name      string         `json:"name" gorm:"not null"`
	Price     float64        `json:"price" gorm:"not null"`
	IsActive  bool           `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	// Pizzas []PizzaToppings `json:"pizzas" gorm:"foreignKey:ToppingID"`
}
