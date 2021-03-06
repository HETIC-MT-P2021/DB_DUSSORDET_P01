package models

import "go_exercise/database"

//Product is a product DTO
type Product struct {
	Code            string               `json:"product_code"`
	Name            string               `json:"product_name"`
	Line            string               `json:"product_line"`
	Scale           string               `json:"product_scale"`
	Vendor          string               `json:"product_vendor"`
	Description     string               `json:"product_description"`
	QuantityInStock int64                `json:"quantity_in_stock"`
	BuyPrice        database.NullFloat64 `json:"buy_price"`
	MSRP            database.NullFloat64 `json:"msrp"`
}
