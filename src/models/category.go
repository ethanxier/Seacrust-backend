package models

import (
	"gorm.io/gorm"
)

const (
	Konsumen       = "Konsumen"
	Tengkulak      = "Tengkulak"
	Pembudidaya    = "Pembudidaya"
	NelayanTangkap = "Nelayan Tangkap"
)

type Category struct {
	gorm.Model
	Name   string   `json:"name" gorm:"type:VARCHAR(255);NOT NULL"`
	Produk []Produk `json:"produk"`
}
