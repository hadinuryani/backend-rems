package dto

type RoleRequest struct {
	RoleName    string  `json:"role_name" binding:"required"`
	Level       *int8   `json:"level" binding:"required"`
	Scope       *string `json:"scope"`
	Description *string `json:"description"`
}

type UpdateRoleRequest struct {
	RoleName    string `json:"role_name" binding:"required"`
	Level       int    `json:"level" binding:"required"`
	Scope       string `json:"scope"`
	Description string `json:"description"`
}

type RoleSalaryRequest struct {
	RoleID     int     `json:"role_id"`
	BaseSalary float64 `json:"base_salary"`
}

type User struct {
	Password   string    
}