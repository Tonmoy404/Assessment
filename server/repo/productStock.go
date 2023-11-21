package repo

import (
	"database/sql"

	"github.com/Tonmoy404/Assessment/service"
)

type ProductStockRepo interface {
	service.ProductStockRepo
}

type productStockRepo struct {
	db        *sql.DB
	tableName string
}

func NewProductStockRepo(db *sql.DB, tableName string) ProductStockRepo {
	return &productStockRepo{
		db:        db,
		tableName: tableName,
	}
}
