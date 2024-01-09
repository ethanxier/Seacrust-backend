package handler

import (
	"net/http"
	"seacrust-backend/src/models"

	"github.com/gin-gonic/gin"
)

func (h *handler) updateDirectCart(ctx *gin.Context) {
	var reqBody struct {
		ProductID uint `json:"product_id"`
		Quantity  uint `json:"quantity"`
	}

	if err := h.BindBody(ctx, &reqBody); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "failed bind body", nil)
		return
	}

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

	var directCardDB models.DirectCard

	if err := h.db.Model(&directCardDB).Where("user_id = ?", userID).First(&directCardDB).Updates(models.DirectCard{
		ProductID: reqBody.ProductID,
		Quantity:  reqBody.Quantity,
	}).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, "error sini", nil)
		return
	}

	h.SuccessResponse(ctx, http.StatusOK, "Update berhasil", nil)
}

func (h *handler) getDirectCard(ctx *gin.Context) {
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

	var directCardDB models.DirectCard
	if err2 := h.db.Model(models.DirectCard{}).Where("user_id = ?", userID).First(&directCardDB).Error; err2 != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, err2.Error(), nil)
		return
	}

	var produkDB models.Produk
	if err := h.db.Model(models.Produk{}).Where("id = ?", directCardDB.ProductID).First(&produkDB).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	var produkResp models.DirectCardResp

	var tokoDB models.Toko
	err := h.db.Where("id = ?", produkDB.TokoID).First(&tokoDB).Error
	if err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var addresses []models.Address
	err3 := h.db.Where("user_id = ?", userID).Find(&addresses).Error
	if err3 != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, err3.Error(), nil)
		return
	}

	produkResp.ID = produkDB.ID
	produkResp.Foto = produkDB.Foto
	produkResp.Name = produkDB.Name
	produkResp.Harga = produkDB.Harga
	produkResp.NamaToko = tokoDB.Name
	produkResp.Stok = produkDB.Stok
	produkResp.Quantity = directCardDB.Quantity
	produkResp.Alamat = addresses

	h.SuccessResponse(ctx, http.StatusOK, "Success", produkResp)
}
