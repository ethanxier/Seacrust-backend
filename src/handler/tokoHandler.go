package handler

import (
	"net/http"
	"seacrust-backend/src/models"

	"github.com/gin-gonic/gin"
)

func (h *handler) tokoRegistrasi(ctx *gin.Context) {
	user, exist := ctx.Get("user")
	if !exist {
		h.ErrorResponse(ctx, http.StatusBadRequest, "Unauthorized", nil)
		return
	}

	claims, ok := user.(models.UserClaims)
	if !ok {
		h.ErrorResponse(ctx, http.StatusBadRequest, "invalid token", nil)
		return
	}

	userID := claims.ID

	var tokoBody models.TokoRegistrasi

	if err := h.BindBody(ctx, &tokoBody); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "Please enter valid data", nil)
		return
	}

	tokoDB := models.Toko{
		Name:      tokoBody.Name,
		Alamat:    tokoBody.Alamat,
		Provinsi:  tokoBody.Provinsi,
		Kota:      tokoBody.Kota,
		Kecamatan: tokoBody.Kecamatan,
		Desa:      tokoBody.Desa,
		KodePos:   tokoBody.KodePos,
		UserID:    userID,
	}

	if err := h.db.Create(&tokoDB).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "Failed to register store", nil)
		return
	}

	h.SuccessResponse(ctx, http.StatusOK, "Store Registration Successful", nil)
}

func (h *handler) getMyToko(ctx *gin.Context) {
	user, exist := ctx.Get("user")
	if !exist {
		h.ErrorResponse(ctx, http.StatusBadRequest, "Unauthorized", nil)
		return
	}

	claims, ok := user.(models.UserClaims)
	if !ok {
		h.ErrorResponse(ctx, http.StatusBadRequest, "invalid token", nil)
		return
	}

	userID := claims.ID

	var tokoDB models.Toko
	err := h.db.Where("user_id = ?", userID).Take(&tokoDB).Error
	if err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var produkDB []models.Produk
	err2 := h.db.Where("toko_id = ?", tokoDB.ID).Take(&produkDB).Error
	if err2 != nil {
		// h.ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		// return
		produkDB = nil
	}

	tokoResp := models.GetToko{
		Name:   tokoDB.Name,
		Produk: produkDB,
		Alamat: tokoDB.Alamat,
	}

	h.SuccessResponse(ctx, http.StatusOK, "Succes", tokoResp)
}
