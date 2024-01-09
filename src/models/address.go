package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	User         User `gorm:"foreignkey:UserID"`
	UserID       uint
	NamaPenerima string `json:"nama_penerima" gorm:"type:VARCHAR(255)"`
	NomorHP      string `json:"nomor_hp"`
	Alamat       string `json:"alamat"`
	Provinsi     string `json:"provinsi"`
	Kota         string `json:"kota"`
	Kecamatan    string `json:"kecamatan"`
	Desa         string `json:"desa"`
	KodePos      string `json:"kode_pos"`
}
