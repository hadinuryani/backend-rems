package handler

import (
	"log"
	"net/http"
	"strconv"

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

func(h *UsersHandler)InsertRoleSalary(c *gin.Context){
	roleId := c.Param("role_id")
	idInt,err := strconv.Atoi(roleId)
	if err != nil{
		log.Println(err.Error())
		response.ErrorResponse(c,http.StatusBadRequest,"id tidak valid",nil)
		log.Println(idInt)
		return
	}
	var req dto.RoleSalaryRequest

	if err := c.ShouldBindJSON(&req);err != nil {
		log.Println(err.Error())
		response.ErrorResponse(c,http.StatusBadRequest,"validation error",nil)
		return
	}
	roleSalary := model.RoleSalary{
		BaseSalary: req.BaseSalary,
	}
	id,err := h.service.InsertRoleSalary(idInt,roleSalary)
	if err != nil {
		log.Println(err.Error())
		response.ErrorResponse(c,http.StatusInternalServerError,"gagal insert role salary",nil)
		return
	}
	response.SuccesResponse(c,http.StatusCreated,"role salary berhasil di insert",gin.H{"id" : id})
}

func(h *UsersHandler)InsertUser(c *gin.Context){
	employeeId := c.Param("employee_id")
	idInt,err := strconv.Atoi(employeeId)
	if err != nil {
		log.Println(err.Error())
		response.ErrorResponse(c,http.StatusBadRequest,"id tidak valid",nil)
		return
	}
	user := model.User{
		Password: "123456",
	}
	id,err := h.service.InsertUser(idInt,user)
	if err != nil {
		log.Println(err.Error())
		response.ErrorResponse(c,http.StatusInternalServerError,"user gagal ter insert",nil)
		return
	}
	
}
