package response

import "github.com/gin-gonic/gin"

type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   any    `json:"error,omitempty"`
}

func SuccesResponse(c *gin.Context,status int,message string,data any) {
	c.JSON(status, BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context,status int,message string,err any){
	c.JSON(status,BaseResponse{
		Success: false,
		Message: message,
		Error: err,
	})
}
