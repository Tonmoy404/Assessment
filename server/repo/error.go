package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Tonmoy404/Assessment/service"
)

type ErrorRepo interface {
	service.ErrorRepo
}

type errorRepo struct {
	db        *sql.DB
	tableName string
}

func NewErrorRepo(db *sql.DB, tableName string) ErrorRepo {
	return &errorRepo{
		db:        db,
		tableName: tableName,
	}
}

func (r *errorRepo) GetError(ctx context.Context, code string) (*service.ErrorDetail, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE code = $1", r.tableName)
	row := r.db.QueryRowContext(ctx, query, code)

	var error service.ErrorDetail
	err := row.Scan(&error.Code, &error.MessageBn)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Return nil if the row does not exist
		}
		return nil, fmt.Errorf("failed to get error: %v", err)
	}

	return &error, nil
}
