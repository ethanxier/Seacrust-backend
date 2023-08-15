package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName     string    `json:"fullname" gorm:"type:VARCHAR(255);NOT NULL"`
	Username     string    `json:"username" gorm:"type:VARCHAR(20);UNIQUE"`
	Email        string    `json:"email" gorm:"type:VARCHAR(255);UNIQUE"`
	Password     string    `json:"password" gorm:"type:VARCHAR(255);NOT NULL"`
	TanggalLahir string    `json:"tanggal_lahir" gorm:"default:null"`
	Domisili     string    `json:"domisili" gorm:"type:VARCHAR(50);default:null"`
	ProfilePhoto string    `json:"profile_photo" gorm:"default:null"`
	Deskripsi    string    `json:"deskripsi" gorm:"type:VARCHAR(250);default:null"`
	NoWhatsapp   string    `json:"no_whatsapp" gorm:"type:VARCHAR(250);default:null"`
	JenisKelamin string    `json:"jenis_kelamin" gorm:"type:VARCHAR(250);default:null"`
	Toko         Toko      `json:"toko"`
	Pesanan      []Pesanan `json:"pesanan"`
}

type UserRegister struct {
	FullName string `json:"fullname" gorm:"NOT NULL" binding:"required,max=50"`
	Username string `json:"username" gorm:"NOT NULL" binding:"required,min=5,max=20"`
	Email    string `json:"email" gorm:"NOT NULL" binding:"required,email"`
	Password string `json:"password" gorm:"NOT NULL" binding:"required,min=8"`
}

type UserLogin struct {
	Email    string `json:"email" gorm:"NOT NULL" binding:"required"`
	Password string `json:"password" gorm:"NOT NULL" binding:"required"`
}

type UserNavbar struct {
	FullName     string `json:"fullname"`
	ProfilePhoto string `json:"profile_photo" gorm:"default:null"`
}

type UserProfilePage struct {
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
