package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName     string     `json:"fullname" gorm:"type:VARCHAR(255);NOT NULL"`
	Username     string     `json:"username" gorm:"type:VARCHAR(20);UNIQUE"`
	Email        string     `json:"email" gorm:"type:VARCHAR(255);UNIQUE"`
	Password     string     `json:"password" gorm:"type:VARCHAR(255);NOT NULL"`
	TanggalLahir string     `json:"tanggal_lahir" gorm:"default:null"`
	Domisili     string     `json:"domisili" gorm:"type:VARCHAR(50);default:null"`
	ProfilePhoto string     `json:"profile_photo" gorm:"default:null"`
	Deskripsi    string     `json:"deskripsi" gorm:"type:VARCHAR(250);default:null"`
	NoWhatsapp   string     `json:"no_whatsapp" gorm:"type:VARCHAR(250);default:null"`
	JenisKelamin string     `json:"jenis_kelamin" gorm:"type:VARCHAR(250);default:null"`
	DirectCard   DirectCard `json:"direct_card"`
	Toko         Toko       `json:"toko"`
	Order        []Order    `json:"order"`
}

type UserRegister struct {
	FullName string `json:"fullname" gorm:"NOT NULL" binding:"required,max=50"`
	Username string `json:"username" gorm:"NOT NULL" binding:"required,min=5,max=20"`
	Email    string `json:"email" gorm:"NOT NULL" binding:"required,email"`
	Password string `json:"password" gorm:"NOT NULL" binding:"required,min=8"`
}

type DirectCard struct {
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

type DirectCardResp struct {
	ID       uint      `json:"id"`
	Foto     string    `json:"foto"`
	Name     string    `json:"name"`
	Harga    float32   `json:"harga"`
	NamaToko string    `json:"nama_toko"`
	Stok     uint      `json:"stok"`
	Quantity uint      `json:"quantity"`
	Alamat   []Address `json:"alamat"`
}

type UserLogin struct {
	Email    string `json:"email" gorm:"NOT NULL" binding:"required"`
	Password string `json:"password" gorm:"NOT NULL" binding:"required"`
}

type UserNavbar struct {
	Username     string `json:"username"`
	ProfilePhoto string `json:"profile_photo" gorm:"default:null"`
}

type UserProfilePage struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	FullName     string `json:"fullname"`
	JenisKelamin string `json:"jenis_kelamin"`
	ProfilePhoto string `json:"profile_photo"`
	Deskripsi    string `json:"deskripsi"`
	TanggalLahir string `json:"tanggal_lahir"`
	Domisili     string `json:"domisili"`
	NoWhatsapp   string `json:"no_whatsapp"`
}

type UserUpdateProfile struct {
	FullName     string `json:"full_name"`
	Domisili     string `json:"domisili"`
	Deskripsi    string `json:"deskripsi"`
	TanggalLahir string `json:"tanggal_lahir"`
	NoWhatsapp   string `json:"no_whatsapp"`
	JenisKelamin string `json:"jenis_kelamin"`
}

type UserAddAddress struct {
	NamaPenerima string `json:"nama_penerima"`
	NomorHP      string `json:"nomor_hp"`
	Alamat       string `json:"alamat"`
	Provinsi     string `json:"provinsi"`
	Kota         string `json:"kota"`
	Kecamatan    string `json:"kecamatan"`
	Desa         string `json:"desa"`
	KodePos      string `json:"kode_pos"`
}

type UserClaims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}

func NewUserClaims(id uint, exp time.Duration) UserClaims {
	return UserClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
		},
	}
}
