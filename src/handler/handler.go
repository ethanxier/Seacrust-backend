package handler

import (
	"fmt"
	"os"

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
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	config.AllowCredentials = true
	rest.http.Use(cors.New(config))

	// Apply CORS middleware to the whole application
	rest.http.Use(cors.New(config))

	rest.registerRoutes()

	return &rest
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
}
