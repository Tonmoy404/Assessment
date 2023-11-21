package repo

import (
	"database/sql"

	"github.com/Tonmoy404/Assessment/service"
)

type CategoryRepo interface {
	service.CategoryRepo
}

type categoryRepo struct {
	db        *sql.DB
	tableName string
}

func NewCategoryRepo(db *sql.DB, tableName string) CategoryRepo {
	return &categoryRepo{
		db:        db,
		tableName: tableName,
	}
}
