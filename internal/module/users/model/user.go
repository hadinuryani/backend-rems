package model

type User struct {
	ID         int        `db:"id"          json:"id"`
	EmployeeID int        `db:"employee_id" json:"employee_id"`
	Password   string     `db:"password"    json:"-"`
	IsActive   bool       `db:"is_active"   json:"is_active"`
}
