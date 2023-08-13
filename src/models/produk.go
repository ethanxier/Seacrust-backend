package models

import (
	"gorm.io/gorm"
)

type Produk struct {
	gorm.Model
	Name       string    `json:"name" gorm:"type:VARCHAR(255);NOT NULL"`
	Deskripsi  string    `json:"deskripsi" gorm:"type:VARCHAR(255);NOT NULL"`
	Foto       string    `json:"foto" gorm:"default:null"`
	Harga      float32   `json:"harga" gorm:"default:null"`
	Stok       int64     `json:"stok" gorm:"default:null"`
	CategoryID uint      `json:"category_id"`
	TokoID     uint      `json:"toko_id"`
	Pesanan    []Pesanan `json:"pesanan" gorm:"foreignkey:ProdukID"`
}

type ProdukResp struct {
	ID        uint    `json:"id"`
	Foto      string  `json:"foto"`
	Name      string  `json:"name"`
	Deskripsi string  `json:"deskripsi"`
	Harga     float32 `json:"harga"`
}
