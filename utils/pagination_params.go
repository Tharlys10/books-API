package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PaginationParams(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.Request.URL.Query().Get("page"))

		if page <= 0 {
			page = 1
		}

		amount, _ := strconv.Atoi(c.Request.URL.Query().Get("amount"))

		switch {
		case amount <= 0:
			amount = 10
			break
		case amount >= 100:
			amount = 100
			break
		}

		offset := (page - 1) * amount

		return db.Offset(offset).Limit(amount)
	}
}
