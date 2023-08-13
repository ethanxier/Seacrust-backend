package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	User         User `gorm:"foreignkey:UserID"`
	UserID       uint
	Namapenerima string `json:"nama_penerima" gorm:"type:VARCHAR(255)"`
	Nomorhp      string `json:"nomor_hp" gorm:"type:VARCHAR(255)"`
	Kota         string `json:"kota" gorm:"type:VARCHAR(255)"`
	Alamat       string `json:"alamat" gorm:"type:VARCHAR(255)"`
}
