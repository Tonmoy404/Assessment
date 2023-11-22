package repo

import (
	"github.com/Tonmoy404/Assessment/service"
	"github.com/jmoiron/sqlx"
)

type SupplierRepo interface {
	service.SupplierRepo
}

type supplierRepo struct {
	db *sqlx.DB
}

func NewSupplierRepo(db *sqlx.DB) CategoryRepo {
	return &categoryRepo{
		db: db,
	}
}
