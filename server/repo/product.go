package repo

import (
	"database/sql"

	"github.com/Tonmoy404/Assessment/service"
)

type ProductRepo interface {
	service.ProductRepo
}

type productRepo struct {
	db        *sql.DB
	tableName string
}

func NewProductRepo(db *sql.DB, tableName string) ProductRepo {
	return &productRepo{
		db:        db,
		tableName: tableName,
	}
}
