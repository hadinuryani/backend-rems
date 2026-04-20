package model

type Regency struct {
	ID          int64  `db:"id"           json:"id"`
	RegencyName string `db:"regency_name" json:"regency_name"`
	ProvinceID  int64  `db:"province_id"  json:"province_id"`
}