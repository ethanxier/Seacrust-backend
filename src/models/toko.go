package models

import (
	"gorm.io/gorm"
)

type Toko struct {
	gorm.Model
	Name      string   `json:"name" gorm:"type:VARCHAR(255);UNIQUE"`
	Produk    []Produk `json:"produk"`
	Alamat    string   `json:"alamat" gorm:"type:VARCHAR(50)"`
	Provinsi  string   `json:"provinsi" gorm:"type:VARCHAR(50)"`
	Kota      string   `json:"kota" gorm:"type:VARCHAR(50)"`
	Kecamatan string   `json:"kecamatan" gorm:"type:VARCHAR(50)"`
	Desa      string   `json:"desa" gorm:"type:VARCHAR(50)"`
	KodePos   string   `json:"kode_pos"`
	IsActive  bool     `json:"isactive" gorm:"default:false"`
	UserID    uint     `json:"user_id"`
}

type TokoRegistrasi struct {
	Name      string `json:"name"`
	Alamat    string `json:"alamat"`
	Provinsi  string `json:"provinsi"`
	Kota      string `json:"kota"`
	Kecamatan string `json:"kecamatan"`
	Desa      string `json:"desa"`
	KodePos   string `json:"kode_pos"`
}

type GetToko struct {
	Name   string   `json:"name" gorm:"type:VARCHAR(255);NOT NULL"`
	Produk []Produk `json:"produk"`
	Alamat string   `json:"alamat" gorm:"type:VARCHAR(20);UNIQUE"`
}
