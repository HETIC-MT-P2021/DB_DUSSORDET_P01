package models

import (
	"database/sql"
	"time"
)

//Order is an order DTO
type Order struct {
	Number         int       `json:"order_number"`
	CreationDate   time.Time `json:"order_date"`
	RequiredDate   time.Time `json:"required_date"`
	ShippedDate    time.Time `json:"shipped_date"`
	Status         string    `json:"status"`
	Comments       string    `json:"comments"`
	CustomerNumber int       `json:"customer_number"`
}

//OrderDetails is an order details DTO
type OrderDetails struct {
	Number          int             `json:"order_number"`
	ProductCode     string          `json:"product_code"`
	QuantityOrdered int             `json:"quantity_ordered"`
	PriceEach       sql.NullFloat64 `json:"price_each"`
	LineNumber      int             `json:"line_number"`
}
