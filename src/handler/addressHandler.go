package handler

import (
	"fmt"
	"net/http"
	"seacrust-backend/src/models"

	"github.com/gin-gonic/gin"
)

func (h *handler) addAddress(ctx *gin.Context) {
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

	var body models.UserAddAddress

	if err := h.BindBody(ctx, &body); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "Please enter valid data", nil)
		return
	}

	fmt.Println("====" + body.NomorHP)

	AddressDB := models.Address{
		UserID:       userID,
		NamaPenerima: body.NamaPenerima,
		NomorHP:      body.NomorHP,
		Alamat:       body.Alamat,
		Provinsi:     body.Provinsi,
		Kota:         body.Kota,
		Kecamatan:    body.Kecamatan,
		Desa:         body.Desa,
		KodePos:      body.KodePos,
	}

	if err := h.db.Create(&AddressDB).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "Failed to add address", nil)
		return
	}

	h.SuccessResponse(ctx, http.StatusOK, "Add Address Successful", nil)
}

func (h *handler) getAllAddressUser(ctx *gin.Context) {
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

	var addressDB []models.Address
	if err := h.db.Where("user_id = ?", userID).Find(&addressDB).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	h.SuccessResponse(ctx, http.StatusOK, "Succes", addressDB)
}
