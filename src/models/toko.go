package models

import (
	"gorm.io/gorm"
)

type Toko struct {
	gorm.Model
	Name   string   `json:"name" gorm:"type:VARCHAR(255);NOT NULL"`
	Produk []Produk `json:"produk"`
	Alamat string   `json:"alamat" gorm:"type:VARCHAR(20);UNIQUE"`
	UserID uint     `json:"user_id"`
}
