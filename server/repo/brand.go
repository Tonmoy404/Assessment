package repo

import "database/sql"

type BrandRepo interface {
	service.BrandRepo
}

type brandRepo struct {
	db        *sql.DB
	tableName string
}

func NewBrandRepo(db *sql.DB, tableName string) BrandRepo {
	return &brandRepo{
		db:        db,
		tableName: tableName,
	}
}
