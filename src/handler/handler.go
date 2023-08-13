package handler

import (
	"fmt"
	"os"
	"seacrust-backend/src/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	http *gin.Engine
	db   *gorm.DB
}

func Init(db *gorm.DB) *handler {
	rest := handler{
		http: gin.New(),
		db:   db,
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

	api.Use(JwtMiddleware())

	h.http.POST("/user/register", h.userRegister)
	h.http.POST("/user/login", h.userLogin)

	api.GET("/profile", h.userGetProfile)
	api.GET("/navbar", h.userGetNavbar)
	api.GET("/user/profile", h.userGetProfile)
	api.PUT("/user/profile/update", h.userUpdateProfile)

	api.GET("/product/:category", h.getProductByCategory)
}
