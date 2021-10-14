package controllers

import (
	"books-api/config/database"
	"books-api/models"
	"books-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func FindBooks(c *gin.Context) {
	db := database.GetDatabase()

	search := "%" + c.Request.URL.Query().Get("search") + "%"

	var books []models.Book
	var count int64

	err := db.Scopes(utils.PaginationParams(c)).Order("name ASC").Find(&books, "name ILIKE ?", search).Count(&count).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot list books" + err.Error(),
		})
		return
	}

	c.JSON(200, map[string]interface{}{"books": books, "count": count})
}

func FindBookByID(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	id := c.Param("id")

	err := db.First(&book, "id = ?", id).Error
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Cannot find book: " + err.Error(),
			"code":    404,
		})
		return
	}

	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Cannot bing JSON: " + err.Error(),
			"code":    400,
		})
		return
	}

	id, _ := uuid.NewUUID()
	book.ID = id.String()

	err = db.Create(&book).Error

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Cannot create book: " + err.Error(),
			"code":    400,
		})
		return
	}

	c.JSON(200, book)
}

func UpdateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Cannot bing JSON: " + err.Error(),
			"code":    400,
		})
		return
	}

	book.ID = c.Param("id")

	err = db.Updates(&book).Error

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Cannot update book: " + err.Error(),
			"code":    400,
		})
		return
	}

	c.JSON(200, book)
}

func DeleteBook(c *gin.Context) {
	db := database.GetDatabase()

	id := c.Param("id")

	var book *models.Book

	err := db.First(&book, "id = ?", id).Error
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Cannot find book: " + err.Error(),
			"code":    404,
		})
		return
	}

	err = db.Delete(&book, "id = ?", id).Error
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Cannot delete book: " + err.Error(),
			"code":    400,
		})
		return
	}

	c.JSON(200, map[string]interface{}{"delete": true})
}
