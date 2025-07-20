package models

import (
	"time"

	"gorm.io/gorm"
)

// Order represents an order in the system
type Order struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	CustomerID  uint           `json:"customer_id" gorm:"not null"`
	OrderDate   time.Time      `json:"order_date" gorm:"not null"`
	TotalAmount float64        `json:"total_amount" gorm:"not null"`
	Tax         float64        `json:"tax" gorm:"not null"`
	OrderStatus string         `json:"order_status" gorm:"not null;default:'pending'"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	// Customer   Customer    `json:"customer" gorm:"foreignKey:CustomerID"`
	// Invoice    Invoice     `json:"invoice" gorm:"foreignKey:OrderID"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
}
