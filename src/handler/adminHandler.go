package handler

import (
	"net/http"
	"os"
	"seacrust-backend/src/jwt"
	"seacrust-backend/src/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h *handler) adminLogin(ctx *gin.Context) {
	var adminBody models.AdminLogin

	if err := h.BindBody(ctx, &adminBody); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "invalid request login", nil)
		return
	}

	var admin models.AdminLogin

	adminUname := os.Getenv("AK")
	if adminBody.Key != adminUname {
		h.ErrorResponse(ctx, http.StatusBadRequest, "invalid request login an admin", nil)
		return
	}

	hashPW, _ := bcrypt.GenerateFromPassword([]byte(os.Getenv("AP")), bcrypt.DefaultCost)
	adminPW := string(hashPW)

	//cek password
	if err := bcrypt.CompareHashAndPassword([]byte(adminPW), []byte(adminBody.Password)); err != nil {
		h.ErrorResponse(ctx, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	tokenJwt, err := jwt.GenerateTokenAdmin(admin)
	if err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, "create token failed", nil)
		return
	}

	h.SuccessResponse(ctx, http.StatusOK, "Login Berhasil", gin.H{
		"tokenA": tokenJwt,
	})
}
