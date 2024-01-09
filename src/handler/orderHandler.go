package handler

import (
	"net/http"
	"seacrust-backend/src/models"

	"github.com/gin-gonic/gin"
)

func (h *handler) getAllUserOrder(ctx *gin.Context) {
	user, exist := ctx.Get("user")
	if !exist {
		h.ErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	claims, ok := user.(models.UserClaims)
	if !ok {
		h.ErrorResponse(ctx, http.StatusBadRequest, "invalid token", nil)
		return
	}

	userID := claims.ID

	var orderDB []models.Order
	err := h.db.Where("user_id = ? AND status != ?", userID, "SELESAI").Find(&orderDB).Error
	if err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	orderResp := []models.OrderResp{}

	for _, order := range orderDB {
		var product models.Produk
		err2 := h.db.Where("id = ?", order.ProdukID).Take(&product).Error
		if err2 != nil {
			h.ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		var tokoDB models.Toko
		err := h.db.Where("id = ?", product.TokoID).Take(&tokoDB).Error
		if err != nil {
			h.ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		orderResp = append(orderResp, models.OrderResp{
			ProdukID:     product.ID,
			ProductPhoto: product.Foto,
			ProductName:  product.Name,
			StoreName:    tokoDB.Name,
			Quantity:     order.Quantity,
			Status:       order.Status,
			TotalCosts:   order.ShippingCosts + int64(product.Harga),
			ProductPrice: product.Harga,
		})
	}

	h.SuccessResponse(ctx, http.StatusOK, "Succes", orderResp)
}

func (h *handler) getAllMyHistoryOrder(ctx *gin.Context) {
	user, exist := ctx.Get("user")
	if !exist {
		h.ErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	claims, ok := user.(models.UserClaims)
	if !ok {
		h.ErrorResponse(ctx, http.StatusBadRequest, "invalid token", nil)
		return
	}

	userID := claims.ID

	var orderDB []models.Order
	err := h.db.Where("user_id = ? AND status = ?", userID, "SELESAI").Find(&orderDB).Error
	if err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	orderResp := []models.OrderResp{}

	for _, order := range orderDB {
		var product models.Produk
		err2 := h.db.Where("id = ?", order.ProdukID).Take(&product).Error
		if err2 != nil {
			h.ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		var tokoDB models.Toko
		err := h.db.Where("id = ?", product.TokoID).Take(&tokoDB).Error
		if err != nil {
			h.ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		orderResp = append(orderResp, models.OrderResp{
			ProdukID:     product.ID,
			ProductPhoto: product.Foto,
			ProductName:  product.Name,
			StoreName:    tokoDB.Name,
			Quantity:     order.Quantity,
			Status:       order.Status,
			TotalCosts:   order.ShippingCosts + int64(product.Harga),
			ProductPrice: product.Harga,
		})
	}

	h.SuccessResponse(ctx, http.StatusOK, "Succes", orderResp)
}
