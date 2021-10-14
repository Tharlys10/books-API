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

func FindBooks(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			app.InternalServerError(errors.New("Internal server error"), c)
			return
		}
	}()

	db := database.GetDatabase()

	search := "%" + c.Request.URL.Query().Get("search") + "%"

	var books []models.Book
	var count int64

	err := db.Scopes(utils.PaginationParams(c)).Order("name ASC").Find(&books, "name ILIKE ?", search).Count(&count).Error
	if err != nil {
		app.BadRequest(errors.New(err.Error()), c)
		return
	}

	app.OK(map[string]interface{}{"books": books, "count": count}, c)
}

func FindBookByID(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			app.InternalServerError(errors.New("Internal server error"), c)
			return
		}
	}()

	db := database.GetDatabase()

	var book models.Book

	id := c.Param("id")

	err := db.First(&book, "id = ?", id).Error
	if err != nil {
		app.BadRequest(errors.New(err.Error()), c)
		return
	}

	app.OK(book, c)
}

func CreateBook(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			app.InternalServerError(errors.New("Internal server error"), c)
			return
		}
	}()

	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		app.BadRequest(errors.New(err.Error()), c)
		return
	}

	id, _ := uuid.NewUUID()
	book.ID = id.String()

	err = db.Create(&book).Error

	if err != nil {
		app.BadRequest(errors.New(err.Error()), c)
		return
	}

	app.OK(book, c)
}

func UpdateBook(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			app.InternalServerError(errors.New("Internal server error"), c)
			return
		}
	}()

	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		app.BadRequest(errors.New(err.Error()), c)
		return
	}

	book.ID = c.Param("id")

	err = db.Updates(&book).Error

	if err != nil {
		app.BadRequest(errors.New(err.Error()), c)
		return
	}

	app.OK(book, c)
}

func DeleteBook(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			app.InternalServerError(errors.New("Internal server error"), c)
			return
		}
	}()

	db := database.GetDatabase()

	id := c.Param("id")

	var book *models.Book

	err := db.First(&book, "id = ?", id).Error
	if err != nil {
		app.NotFound(errors.New(err.Error()), c)
		return
	}

	err = db.Delete(&book, "id = ?", id).Error
	if err != nil {
		app.BadRequest(errors.New(err.Error()), c)
		return
	}

	app.OK(map[string]interface{}{"delete": true}, c)
}
