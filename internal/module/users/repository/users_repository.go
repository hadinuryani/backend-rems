package repository

import (
	"database/sql"

	"github.com/backend-rems/internal/module/users/model"
)

type UsersRepository interface {
	InsertRole(role model.Role)(int64,error)
	InsertRoleSalary(roleId int,roleSalary model.RoleSalary)(int64,error)
}

type usersRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB)UsersRepository{
	return &usersRepository{db}
}

func(r *usersRepository)InsertRole(role model.Role)(int64,error){
	query := `INSERT INTO roles 
				(role_name,level,scope,description,is_active,created_at)
				VALUES(?,?,?,?,?,NOW())`
	result,err := r.db.Exec(query,role.RoleName,role.Level,role.Scope,role.Description,true)
	if err != nil {
		return 0,err
	}
	id,err := result.LastInsertId()
	if err != nil {
		return 0,err
	}
	return id,nil
}

func(r *usersRepository)InsertRoleSalary(roleId int,roleSalary model.RoleSalary)(int64,error){
	query := `INSERT INTO role_salaries 
				(role_id,base_salary,created_at)
				VALUES(?,?,NOW())`
	
	result,err := r.db.Exec(query,roleId,roleSalary.BaseSalary)
	if err != nil {
		return 0,err
	}
	id,err := result.LastInsertId()
	if err != nil {
		return 0,nil
	}
	return id,nil
}


