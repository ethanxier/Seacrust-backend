package models

import (
	"gorm.io/gorm"
)

type Produk struct {
	gorm.Model
	Name       string  `json:"name" gorm:"type:VARCHAR(255);NOT NULL"`
	Deskripsi  string  `json:"deskripsi" gorm:"type:VARCHAR(255);NOT NULL"`
	Foto       string  `json:"foto" gorm:"default:null"`
	Harga      float32 `json:"harga" gorm:"default:null"`
	Stok       uint    `json:"stok" gorm:"default:null"`
	Terjual    uint    `json:"terjual"`
	CategoryID uint    `json:"category_id"`
	TokoID     uint    `json:"toko_id"`
	Order      []Order `json:"order" gorm:"foreignkey:ProdukID"`
	IsVerified bool    `json:"is_verified"`
}

type ProdukResp struct {
	ID           uint    `json:"id"`
	Foto         string  `json:"foto"`
	Name         string  `json:"name"`
	Deskripsi    string  `json:"deskripsi"`
	Harga        float32 `json:"harga"`
	IsVerified   bool    `json:"is_verified"`
	NamaToko     string  `json:"nama_toko"`
	Kategori     string  `json:"kategori"`
	DomisiliToko string  `json:"domisili_toko"`
}

type ProdukDetailResp struct {
	ID        uint    `json:"id"`
	Foto      string  `json:"foto"`
	Name      string  `json:"name"`
	Deskripsi string  `json:"deskripsi"`
	Harga     float32 `json:"harga"`
	NamaToko  string  `json:"nama_toko"`
	Terjual   uint    `json:"terjual"`
	Stok      uint    `json:"stok"`
}

type NewProduct struct {
	NamaProduk string  `json:"nama_produk"`
	Deskripsi  string  `json:"deskripsi"`
	Harga      float32 `json:"harga"`
	Stok       uint    `json:"stok"`
	KategoriID uint    `json:"kategori"`
}
