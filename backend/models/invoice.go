package models

import (
	"time"

	"gorm.io/gorm"
)

// Invoice represents an invoice in the system
type Invoice struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	OrderID        uint           `json:"order_id" gorm:"not null;uniqueIndex"`
	InvoiceNumber  string         `json:"invoice_number" gorm:"unique;not null"`
	InvoiceDate    time.Time      `json:"invoice_date" gorm:"not null"`
	SubtotalAmount float64        `json:"subtotal_amount" gorm:"not null"`
	TaxAmount      float64        `json:"tax_amount" gorm:"not null"`
	TotalAmount    float64        `json:"total_amount" gorm:"not null"`
	PaymentStatus  string         `json:"payment_status" gorm:"not null;default:'pending'"` // pending, paid, overdue, cancelled
	PaymentDate    *time.Time     `json:"payment_date"`
	Notes          string         `json:"notes"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	Order Order `json:"order" gorm:"foreignKey:OrderID"`
}
