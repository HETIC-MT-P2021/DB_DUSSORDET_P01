package models

import (
	"go_exercise/database"
	"time"
)

//Order is an order DTO
type Order struct {
	Number         int                 `json:"order_number"`
	CreationDate   time.Time           `json:"order_date"`
	RequiredDate   time.Time           `json:"required_date"`
	ShippedDate    database.NullTime   `json:"shipped_date"`
	Status         string              `json:"status"`
	Comments       database.NullString `json:"comments"`
	CustomerNumber int                 `json:"customer_number"` //@toDo maybe to remove?
	NumberOfItems  int                 `json:"number_of_items"`
	TotalCost      float32             `json:"total_cost"`
}

////OrderDetails is an order details DTO
//type OrderDetails struct {
//	Number          int                  `json:"order_number"`
//	ProductCode     string               `json:"product_code"`
//	QuantityOrdered int                  `json:"quantity_ordered"`
//	PriceEach       database.NullFloat64 `json:"price_each"`
//	LineNumber      int                  `json:"line_number"`
//}
