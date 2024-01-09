package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ProdukID      uint   `json:"produk_id"`
	Produk        Produk `json:"produk"`
	UserID        uint   `json:"user_id"`
	Quantity      uint   `json:"quantity"`
	Status        string `json:"status"`
	ShippingCosts int64  `json:"shipping_costs" gorm:"default:0"`
}

type OrderResp struct {
	ProductName  string  `json:"product_name"`
	ProductPhoto string  `json:"product_photo"`
	ProdukID     uint    `json:"produk_id"`
	ProductPrice float32 `json:"product_price"`
	StoreName    string  `json:"store_name"`
	Quantity     uint    `json:"quantity"`
	Status       string  `json:"status"`
	TotalCosts   int64   `json:"total_cost"`
}
