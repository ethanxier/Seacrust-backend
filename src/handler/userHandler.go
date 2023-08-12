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
			h.ErrorResponse(ctx, http.StatusBadRequest, "invalid username or password", nil)
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

	ctx.Header("Authorization", "Bearer "+tokenJwt)
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
		FullName:     userDB.FullName,
		ProfilePhoto: "https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50",
	}

	h.SuccessResponse(ctx, http.StatusOK, "Succes", userRes)
}
