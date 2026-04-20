package model

import "time"

type RoleSalary struct {
	ID         int        `db:"id"         json:"id"`
	RoleID     int        `db:"role_id"    json:"role_id"`
	BaseSalary float64    `db:"base_salary" json:"base_salary"`
	UpdatedAt  *time.Time `db:"updated_at" json:"updated_at"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
}
