package repo

import (
	"github.com/Tonmoy404/Assessment/service"
	"github.com/jmoiron/sqlx"
)

type CategoryRepo interface {
	service.CategoryRepo
}

type categoryRepo struct {
	db *sqlx.DB
}

func NewCategoryRepo(db *sqlx.DB) CategoryRepo {
	return &categoryRepo{
		db: db,
	}
}
