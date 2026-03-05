package shared

import "github.com/gin-gonic/gin"

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   any    `json:"error,omitempty"`
}

func SuccessResponse(ctx *gin.Context, statusCode int, message string, data any) {
	ctx.JSON(statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
		Error:   nil,
	})
}

func ErrorResponse(ctx *gin.Context, statusCode int, message string, err any) {
	ctx.JSON(statusCode, Response{
		Success: false,
		Message: message,
		Data:    nil,
		Error:   err,
	})
}