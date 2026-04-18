package model

type TransactionCategory struct {
	ID   int    `db:"id"   json:"id"`
	Name string `db:"name" json:"name"`
	Type string `db:"type" json:"type"` 
}
