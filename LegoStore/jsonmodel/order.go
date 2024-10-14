package jsonmodel

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderId     uint32     `json:"order_id"`
	CustomerId  uint32     `json:"customer_id"`
	LineItems   []LineItem `json:"line_items"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	ShippedAt   *time.Time `json:"shipped_at,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

type LineItem struct {
	ItemId   uuid.UUID `json:"item_id"`
	Quantity uint      `json:"quantity"`
	Price    uint      `json:"price"`
}
