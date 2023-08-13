package handler

import (
	"net/http"
	"seacrust-backend/src/models"

	"github.com/gin-gonic/gin"
)

func (h *handler) getProductByCategory(ctx *gin.Context) {
	var kategori string
	if err := h.BindParam(ctx, &kategori); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "failed to bind body", nil)
		return
	}

	var kategoriID uint
	switch kategori {
	case "konsumen":
		kategoriID = 1
	case "tengkulak":
		kategoriID = 2
	case "pembudidaya":
		kategoriID = 3
	case "nelayan tangkap":
		kategoriID = 4
	default:
		kategoriID = 0
	}

	var produkDB []models.Produk

	db := h.db.Model(models.Produk{})

	if kategoriID != 0 {
		db = db.Where("kategori_id = ?", kategoriID)
	}

	if err := db.Find(&produkDB).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	var produkResp []models.ProdukResp
	for _, prd := range produkDB {
		var produk models.ProdukResp

		produk.ID = prd.ID
		produk.Foto = prd.Foto
		produk.Name = prd.Name
		produk.Deskripsi = prd.Deskripsi
		produk.Harga = prd.Harga

	}

	h.SuccessResponse(ctx, http.StatusOK, "Success", produkResp)
}
