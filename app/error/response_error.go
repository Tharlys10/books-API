package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// errorMessage interface definition response json
type errorMessage struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

//UnprocessableEntity return response for o status HTTP 422
func UnprocessableEntity(err error, c *gin.Context) {
	data := errorMessage{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    err.Error(),
	}

	c.JSON(data.StatusCode, data)
}

//BadRequest return response for o status HTTP 400
func BadRequest(err error, c *gin.Context) {
	data := errorMessage{
		StatusCode: http.StatusBadRequest,
		Message:    err.Error(),
	}

	c.JSON(data.StatusCode, data)
}

//RequestEntityTooLarge return response for o status HTTP 413
func RequestEntityTooLarge(err error, c *gin.Context) {
	data := errorMessage{
		StatusCode: http.StatusRequestEntityTooLarge,
		Message:    err.Error(),
	}

	c.JSON(data.StatusCode, data)
}

//Conflict return response for o status HTTP 409
func Conflict(err error, c *gin.Context) {
	data := errorMessage{
		StatusCode: http.StatusConflict,
		Message:    err.Error(),
	}

	c.JSON(data.StatusCode, data)
}

//InternalServerError return response for o status HTTP 500
func InternalServerError(err error, c *gin.Context) {
	data := errorMessage{
		StatusCode: http.StatusInternalServerError,
		Message:    err.Error(),
	}

	c.JSON(data.StatusCode, data)
}

//NotFound return response for o status HTTP 404
func NotFound(err error, c *gin.Context) {
	data := errorMessage{
		StatusCode: http.StatusNotFound,
		Message:    err.Error(),
	}

	c.JSON(data.StatusCode, data)
}

//Forbidden return response for o status HTTP 403
func Forbidden(err error, c *gin.Context) {
	data := errorMessage{
		StatusCode: http.StatusForbidden,
		Message:    err.Error(),
	}

	c.JSON(data.StatusCode, data)
}

//Unauthorized return response for o status HTTP 401
func Unauthorized(err error, c *gin.Context) {
	data := errorMessage{
		StatusCode: 401,
		Message:    err.Error(),
	}

	c.JSON(data.StatusCode, data)
}

//OK return uma response for status HTTP 200
func OK(data interface{}, c *gin.Context) {
	c.JSON(200, data)
}
