package controllers

import (
	app "books-api/app/error"
	"books-api/config/database"
	"books-api/models"
	"books-api/utils"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateUser(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			app.InternalServerError(errors.New("Internal server error"), ctx)
			return
		}
	}()

	db := database.GetDatabase()

	var payload models.User

	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		app.BadRequest(errors.New(err.Error()), ctx)
		return
	}

	var user models.User

	_ = db.First(&user, "email = ?", payload.Email)

	if user.ID != "" {
		app.BadRequest(errors.New("User already exists"), ctx)
		return
	}

	id, _ := uuid.NewUUID()
	payload.ID = id.String()

	payload.Password = utils.SHA256Encoder(payload.Password)

	err = db.Create(&payload).Error

	if err != nil {
		app.BadRequest(errors.New(err.Error()), ctx)
		return
	}

	userWithOutPassword := models.UserWithOutPassword{
		ID:        payload.ID,
		Name:      payload.Name,
		Email:     payload.Email,
		CreatedAt: payload.CreatedAt,
		UpdatedAt: payload.UpdatedAt,
		DeletedAt: payload.DeletedAt,
	}

	app.OK(userWithOutPassword, ctx)
}
