package model

type Province struct {
	ID           int64  `db:"id"            json:"id"`
	ProvinceName string `db:"province_name" json:"province_name"`
}