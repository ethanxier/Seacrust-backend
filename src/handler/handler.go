package handler

import (
	"fmt"
	"os"
	"seacrust-backend/src/models"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	http      *gin.Engine
	db        *gorm.DB
	supClient supabasestorageuploader.SupabaseClientService
}

func Init(db *gorm.DB, supClient supabasestorageuploader.SupabaseClientService) *handler {
	rest := handler{
		http:      gin.New(),
		db:        db,
		supClient: supClient,
	}

	// CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "OPTIONS", "PUT"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	config.AllowCredentials = true
	rest.http.Use(cors.New(config))

	// Apply CORS middleware to the whole application
	rest.http.Use(cors.New(config))

	rest.registerRoutes()

	return &rest
}

func (h *handler) SeedCategory(sql *gorm.DB) error {
	var categories []models.Category

	if err := sql.First(&categories).Error; err != gorm.ErrRecordNotFound {
		return err
	}
	categories = []models.Category{
		{
			Name: models.Konsumen,
		},
		{
			Name: models.Tengkulak,
		},
		{
			Name: models.Pembudidaya,
		},
		{
			Name: models.NelayanTangkap,
		},
	}

	if err := sql.Create(&categories).Error; err != nil {
		return err
	}
	return nil
}

func (h *handler) Run() {
	h.http.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}

func (h *handler) registerRoutes() {
	api := h.http.Group("/api")
	admin := h.http.Group("/admin")

	api.Use(JwtMiddleware())
	admin.Use(JwtMiddlewareAdmin())

	h.http.POST("/admin/login", h.adminLogin)
	admin.GET("verif/toko", h.getAllUnverifiedToko)
	admin.GET("verif/produk", h.getAllUnverifiedProduct)
	admin.PUT("verif/toko/:id", h.setVerifToko)
	admin.PUT("verif/produk/:id", h.setVerifProduk)

	h.http.POST("/user/register", h.userRegister)
	h.http.POST("/user/login", h.userLogin)

	api.GET("/profile", h.userGetProfile)
	api.GET("/navbar", h.userGetNavbar)
	api.GET("/user/profile", h.userGetProfile)
	api.PUT("/user/profile/update", h.userUpdateProfile)
	api.PUT("/user/profile/update/photo", h.userUpdatePhotoProfile)
	api.GET("/user/toko", h.getMyToko)
	api.POST("/user/toko/regis", h.tokoRegistrasi)
	api.POST("/user/toko/create-product", h.createProduct)
	api.GET("user/my-order", h.getAllUserOrder)
	api.GET("user/my-history", h.getAllMyHistoryOrder)

	h.http.GET("/produk/:category_id", h.getAllProductByCategory)
	h.http.GET("/produk/search/:key", h.getAllProductBySearchKey)
	api.GET("/produk/detail/:id", h.getProductByID)
	api.PUT("/produk/direct-cart", h.updateDirectCart)
	api.GET("/produk/direct-cart", h.getDirectCard)
	// api.POST("/produk/checkout/:id", h.checkOutProduk)

	api.POST("/user/address/add", h.addAddress)
	api.GET("user/address/get", h.getAllAddressUser)

}
