package repo

import (
	"database/sql"

	"github.com/Tonmoy404/Assessment/service"
)

type SupplierRepo interface {
	service.SupplierRepo
}

type supplierRepo struct {
	db        *sql.DB
	tableName string
}

func NewSupplierRepo(db *sql.DB, tableName string) CategoryRepo {
	return &categoryRepo{
		db:        db,
		tableName: tableName,
	}
}
