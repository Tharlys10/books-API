package middlewares

import (
	app "books-api/app/error"
	"books-api/utils"
	"errors"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const Bearer_schema = "Bearer "
		header := ctx.GetHeader("Authorization")
		if header == "" {
			ctx.AbortWithStatus(401)
			app.Unauthorized(errors.New("Not authorized"), ctx)
			return
		}

		token := header[len(Bearer_schema):]

		if !utils.NewJWTService().ValidateToken(token) {
			ctx.AbortWithStatus(401)
			app.Unauthorized(errors.New("Not authorized"), ctx)
			return
		}
	}
}
