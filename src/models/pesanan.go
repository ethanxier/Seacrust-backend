package models

import "gorm.io/gorm"

type Pesanan struct {
	gorm.Model
	ProdukID      uint   `json:"produk_id"`
	Produk        Produk `json:"produk"`
	UserID        uint   `json:"user_id"`
	Quantity      uint   `json:"quantity"`
	Totalprice    uint   `json:"totalprice"`
	Status        string `json:"status"`
	ShippingCosts int64  `json:"shipping_costs" gorm:"default:0"`
}
