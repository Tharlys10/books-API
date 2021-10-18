package controllers

import (
	app "books-api/app/error"
	"books-api/config/database"
	"books-api/models"
	"books-api/types"
	"books-api/utils"
	"errors"

	"github.com/gin-gonic/gin"
)

func SendLogin(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			app.InternalServerError(errors.New("Internal server error"), ctx)
			return
		}
	}()

	db := database.GetDatabase()

	var payload types.LoginPayload

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		app.BadRequest(errors.New("Payload missing from login: "+err.Error()), ctx)
		return
	}

	var user models.User

	err = db.First(&user, "email = ?", payload.Email).Error
	if err != nil {
		app.BadRequest(errors.New("E-mail ou password incorrect"), ctx)
		return
	}

	if utils.CompareHashAndPassword([]byte(payload.Password), []byte(user.Password)) {
		app.BadRequest(errors.New("E-mail ou password incorrect"), ctx)
		return
	}

	token, err := utils.NewJWTService().GenerateToken(user.ID)
	if err != nil {
		app.InternalServerError(errors.New("Internal server error"), ctx)
		return
	}

	userWithOutPassword := models.UserWithOutPassword{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}

	app.OK(map[string]interface{}{"user": userWithOutPassword, "token": token}, ctx)
}
