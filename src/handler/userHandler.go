package handler

import (
	"fmt"
	"net/http"
	"seacrust-backend/src/jwt"
	"seacrust-backend/src/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h *handler) userRegister(ctx *gin.Context) {
	var userBody models.UserRegister

	if err := h.BindBody(ctx, &userBody); err != nil {
		if len(userBody.Username) < 4 {
			h.ErrorResponse(ctx, http.StatusBadRequest, "Username must be at least 4 characters", nil)
			return
		}
		if len(userBody.Username) > 20 {
			h.ErrorResponse(ctx, http.StatusBadRequest, "Username can be at most 20 characters", nil)
			return
		}
		if len(userBody.Password) < 8 {
			h.ErrorResponse(ctx, http.StatusBadRequest, "Password must be at least 8 characters", nil)
			return
		}
		h.ErrorResponse(ctx, http.StatusBadRequest, "Please enter valid data", nil)
		return
	}

	var userDB models.User

	userDB.FullName = userBody.FullName
	userDB.Username = userBody.Username
	userDB.Email = userBody.Email

	hashPW, _ := bcrypt.GenerateFromPassword([]byte(userBody.Password), bcrypt.DefaultCost)
	userDB.Password = string(hashPW)

	if err := h.db.Create(&userDB).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, "Username or Email has already been used", nil)
		fmt.Println(err)
		return
	}

	var directCardDB models.DirectCard
	directCardDB.UserID = userDB.ID

	if err := h.db.Create(&directCardDB).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, "Username or Email has already been used", nil)
		fmt.Println(err)
		return
	}

	h.SuccessResponse(ctx, http.StatusOK, "Registration successful", nil)
}

func (h *handler) userLogin(ctx *gin.Context) {
	var userBody models.UserLogin

	if err := h.BindBody(ctx, &userBody); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "Please enter valid data", nil)
		return
	}

	var userDB models.User

	if err := h.db.Where("email = ?", userBody.Email).First(&userDB).Error; err != nil {

		if err := h.db.Where("username = ?", userBody.Email).First(&userDB).Error; err != nil {
			h.ErrorResponse(ctx, http.StatusBadRequest, "account notfound", nil)
			return
		}

	}

	//cek password
	if err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(userBody.Password)); err != nil {
		h.ErrorResponse(ctx, http.StatusUnauthorized, "invalid username or password", nil)
		return
	}

	tokenJwt, err := jwt.GenerateToken(userDB)
	if err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, "create token failed", nil)
		return
	}

	h.SuccessResponse(ctx, http.StatusOK, "Login Successfull", gin.H{
		"token": tokenJwt,
	})
}

func (h *handler) userGetProfile(ctx *gin.Context) {
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

	var userDB models.User
	err := h.db.Where("id = ?", userID).Take(&userDB).Error
	if err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	userResp := models.UserProfilePage{
		Username:     userDB.Username,
		Email:        userDB.Email,
		FullName:     userDB.FullName,
		ProfilePhoto: userDB.ProfilePhoto,
		Deskripsi:    userDB.Deskripsi,
		TanggalLahir: userDB.TanggalLahir,
		Domisili:     userDB.Domisili,
		NoWhatsapp:   userDB.NoWhatsapp,
		JenisKelamin: userDB.JenisKelamin,
	}

	h.SuccessResponse(ctx, http.StatusOK, "Succes", userResp)
}

func (h *handler) userGetNavbar(ctx *gin.Context) {
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

	var userDB models.User
	err := h.db.Where("id = ?", userID).Take(&userDB).Error
	if err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	userRes := models.UserNavbar{
		Username:     userDB.Username,
		ProfilePhoto: userDB.ProfilePhoto,
	}

	h.SuccessResponse(ctx, http.StatusOK, "Succes", userRes)
}

func (h *handler) userUpdateProfile(ctx *gin.Context) {
	var userBody models.UserUpdateProfile
	if err := h.BindBody(ctx, &userBody); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "data tidak diterima", nil)
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

	var userDB models.User

	if err := h.db.Model(&userDB).Where("id = ?", userID).First(&userDB).Updates(models.User{
		FullName:     userBody.FullName,
		Domisili:     userBody.Domisili,
		TanggalLahir: userBody.TanggalLahir,
		JenisKelamin: userBody.JenisKelamin,
		Deskripsi:    userBody.Deskripsi,
		NoWhatsapp:   userBody.NoWhatsapp,
	}).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, "error sini", nil)
		return
	}

	h.SuccessResponse(ctx, http.StatusOK, "Update berhasil", nil)
}

func (h *handler) userUpdatePhotoProfile(ctx *gin.Context) {
	file, err := ctx.FormFile("foto")
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
		h.ErrorResponse(ctx, http.StatusBadRequest, "Unauthorized", nil)
		return
	}

	claims, ok := user.(models.UserClaims)
	if !ok {
		h.ErrorResponse(ctx, http.StatusBadRequest, "invalid token", nil)
		return
	}

	userID := claims.ID

	var userDB models.User

	if err := h.db.Model(&userDB).Where("id = ?", userID).First(&userDB).Updates(models.User{
		ProfilePhoto: link,
	}).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, "error sini", nil)
		return
	}

	h.SuccessResponse(ctx, http.StatusOK, "Update berhasil", nil)
}
