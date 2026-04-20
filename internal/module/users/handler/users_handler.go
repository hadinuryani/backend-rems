package handler

import (
	"log"
	"net/http"

	"github.com/backend-rems/internal/module/users/dto"
	"github.com/backend-rems/internal/module/users/model"
	"github.com/backend-rems/internal/module/users/service"
	"github.com/backend-rems/pkg/response"
	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	service service.UsersService
}

func NewUsersHandler(s service.UsersService)*UsersHandler{
	return &UsersHandler{
		service: s,
	}
}

func(h *UsersHandler)InsertRole(c *gin.Context){
	var req dto.RoleRequest
	if err := c.ShouldBindJSON(&req);err != nil{
		response.ErrorResponse(c,http.StatusBadRequest,"validator error",err.Error())
		return
	}
	role := model.Role{
		RoleName: req.RoleName,
		Level: req.Level,
		Scope: req.Scope,
		Description: req.Description,
	}
	id,err := h.service.InsertRole(role)
	if err != nil {
		log.Println(err.Error())
		response.ErrorResponse(c,http.StatusInternalServerError,"gagal insert role",nil)
		return
	}
	response.SuccesResponse(c,http.StatusCreated,"role berhasil di tambahkan",gin.H{"id" : id})
}