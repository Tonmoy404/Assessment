package repo

import (
	"github.com/Tonmoy404/Assessment/service"
	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	service.ProductRepo
}

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}
