package service

import (
	"github.com/backend-rems/internal/module/users/model"
	"github.com/backend-rems/internal/module/users/repository"
)

type UsersService interface {
	InsertRole(role model.Role)(int64,error)
	InsertRoleSalary(roleId int,roleSalary model.RoleSalary)(int64,error)
	InsertUser(employeeId int,user model.User)(int64,error)
}

type usersService struct {
	repo repository.UsersRepository
}

func NewUsersService(repo repository.UsersRepository)UsersService{
	return &usersService{repo}
}

func(s *usersService)InsertRole(role model.Role)(int64,error){
	return s.repo.InsertRole(role)
}

func(s *usersService)InsertRoleSalary(roleId int,roleSalary model.RoleSalary)(int64,error){
	return s.repo.InsertRoleSalary(roleId,roleSalary)
}

func(s *usersService)InsertUser(employeeId int,user model.User)(int64,error){
	return s.repo.InsertUser(employeeId,user)
}

