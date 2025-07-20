package models

import (
	"time"

	"gorm.io/gorm"
)

// OrderItem represents items in an order
type OrderItem struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	OrderID    uint           `json:"order_id" gorm:"not null"`
	ItemID     uint           `json:"item_id" gorm:"not null"`
	Quantity   int            `json:"quantity" gorm:"not null"`
	TotalPrice float64        `json:"total_price" gorm:"not null"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	Order Order `json:"order" gorm:"foreignKey:OrderID"`
	Item  Item  `json:"item" gorm:"foreignKey:ItemID"`
}
