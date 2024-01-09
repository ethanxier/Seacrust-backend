package handler

import (
	"net/http"
	"seacrust-backend/src/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) getAllProductByCategory(ctx *gin.Context) {
	var kategori struct {
		ID uint `uri:"category_id"`
	}
	if err := h.BindParam(ctx, &kategori); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "failed to bind body", nil)
		return
	}

	var produkDB []models.Produk

	db := h.db.Model(models.Produk{})

	if kategori.ID > 0 {
		_ = db.Where("category_id = ?", kategori.ID)
	}

	if err := db.Where("is_verified = ?", true).Find(&produkDB).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	var produkResp []models.ProdukResp
	for _, prd := range produkDB {
		var produk models.ProdukResp

		var tokoDB models.Toko
		err := h.db.Where("id = ?", prd.TokoID).Take(&tokoDB).Error
		if err != nil {
			h.ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		produk.ID = prd.ID
		produk.Foto = prd.Foto
		produk.Name = prd.Name
		produk.Deskripsi = prd.Deskripsi
		produk.Harga = prd.Harga
		produk.DomisiliToko = tokoDB.Kota

		produkResp = append(produkResp, produk)
	}

	h.SuccessResponse(ctx, http.StatusOK, "Success", produkResp)
}

func (h *handler) createProduct(ctx *gin.Context) {
	file, err := ctx.FormFile("foto_produk")
	if err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	link, err := h.supClient.Upload(file)
	if err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

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

	var tokoDB models.Toko
	err2 := h.db.Where("user_id = ?", userID).Take(&tokoDB).Error
	if err2 != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, err2.Error(), nil)
		return
	}

	namaProduk := ctx.PostForm("nama_produk")
	deskripsi := ctx.PostForm("deskripsi")
	stokStr := ctx.PostForm("stok")
	hargaStr := ctx.PostForm("harga")
	kategoriStr := ctx.PostForm("kategori")

	stok, _ := strconv.ParseUint(stokStr, 10, 64)
	harga, _ := strconv.ParseFloat(hargaStr, 32)
	kategori, _ := strconv.ParseUint(kategoriStr, 10, 64)

	productDB := models.Produk{
		Name:       namaProduk,
		Deskripsi:  deskripsi,
		Foto:       link,
		Stok:       uint(stok),
		Harga:      float32(harga),
		CategoryID: uint(kategori),
		TokoID:     tokoDB.ID,
	}

	if err := h.db.Create(&productDB).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "Failed to create product", nil)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
		"data":    productDB, // You can adjust the response as needed
	})
}

func (h *handler) getAllUnverifiedProduct(ctx *gin.Context) {
	var produkDB []models.Produk
	err := h.db.Where("is_verified = ?", false).Find(&produkDB).Error
	if err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	produkResp := []models.ProdukResp{}

	for _, produk := range produkDB {
		var tokoDB models.Toko
		err := h.db.Where("id = ?", produk.TokoID).Take(&tokoDB).Error
		if err != nil {
			h.ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		type category struct {
			ID   uint   `json:"id"`
			Name string `json:"name"`
		}
		var kategoriDB category
		err2 := h.db.Where("id = ?", produk.CategoryID).Take(&kategoriDB).Error
		if err2 != nil {
			h.ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		produkResp = append(produkResp, models.ProdukResp{
			ID:         produk.ID,
			Foto:       produk.Foto,
			Name:       produk.Name,
			Deskripsi:  produk.Deskripsi,
			Harga:      produk.Harga,
			IsVerified: produk.IsVerified,
			NamaToko:   tokoDB.Name,
			Kategori:   kategoriDB.Name,
		})
	}

	h.SuccessResponse(ctx, http.StatusOK, "Succes", produkResp)
}

func (h *handler) setVerifProduk(ctx *gin.Context) {
	var IDParam struct {
		ID uint `uri:"id"`
	}

	if err := h.BindParam(ctx, &IDParam); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "failed to bind param", nil)
		return
	}

	produkDB := models.Produk{}
	if err := h.db.Model(&produkDB).Where("id = ?", IDParam.ID).First(&produkDB).Updates(models.Produk{
		IsVerified: true,
	}).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, "error sini", nil)
		return
	}

	h.SuccessResponse(ctx, http.StatusOK, "Succes", nil)
}

func (h *handler) getProductByID(ctx *gin.Context) {
	var produkReq struct {
		ID uint `uri:"id"`
	}

	if err := h.BindParam(ctx, &produkReq); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "failed to bind param", nil)
		return
	}

	var produkDB models.Produk

	db := h.db.Model(models.Produk{})

	if produkReq.ID > 0 {
		_ = db.Where("id = ?", produkReq.ID)
	}

	if err := db.Where("is_verified = ?", true).First(&produkDB).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	var produkResp models.ProdukDetailResp

	var tokoDB models.Toko
	err := h.db.Where("id = ?", produkDB.TokoID).First(&tokoDB).Error
	if err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	produkResp.ID = produkDB.ID
	produkResp.Foto = produkDB.Foto
	produkResp.Name = produkDB.Name
	produkResp.Deskripsi = produkDB.Deskripsi
	produkResp.Harga = produkDB.Harga
	produkResp.NamaToko = tokoDB.Name
	produkResp.Terjual = produkDB.Terjual
	produkResp.Stok = produkDB.Stok

	h.SuccessResponse(ctx, http.StatusOK, "Success", produkResp)
}

func (h *handler) getAllProductBySearchKey(ctx *gin.Context) {
	var KeyWord struct {
		Key string `uri:"key"`
	}
	if err := h.BindParam(ctx, &KeyWord); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "failed to bind route parameter", nil)
		return
	}

	var produkDB []models.Produk

	db := h.db.Model(models.Produk{})

	if KeyWord.Key != "" {
		db = db.Where("name LIKE ?", "%"+KeyWord.Key+"%")
	}

	if err := db.Where("is_verified = ?", true).Find(&produkDB).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	var produkResp []models.ProdukResp
	for _, prd := range produkDB {
		var produk models.ProdukResp

		var tokoDB models.Toko
		err := h.db.Where("id = ?", prd.TokoID).Take(&tokoDB).Error
		if err != nil {
			h.ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		produk.ID = prd.ID
		produk.Foto = prd.Foto
		produk.Name = prd.Name
		produk.Deskripsi = prd.Deskripsi
		produk.Harga = prd.Harga
		produk.DomisiliToko = tokoDB.Kota

		produkResp = append(produkResp, produk)
	}

	h.SuccessResponse(ctx, http.StatusOK, "Success", produkResp)
}
