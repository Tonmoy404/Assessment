package repo

import (
	"github.com/Tonmoy404/Assessment/service"
	"github.com/jmoiron/sqlx"
)

type BrandRepo interface {
	service.BrandRepo
}

type brandRepo struct {
	db *sqlx.DB
}

func NewBrandRepo(db *sqlx.DB) BrandRepo {
	return &brandRepo{
		db: db,
	}
}
