package repo

import (
	"context"

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

func (r *brandRepo) Create(ctx context.Context, brand *service.Brand) (*service.Brand, error) {
	var newBrand service.Brand
	err := r.db.QueryRowContext(ctx,
		"INSERT INTO brands (id, name, status_id, created_at) VALUES ($1, $2, $3, $4) RETURNING id, name, status_id, created_at",
		brand.Id, brand.Name, brand.StatusId, brand.CreatedAt,
	).Scan(&newBrand.Id, &newBrand.Name, &newBrand.StatusId, &newBrand.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &newBrand, nil
}

func (r *brandRepo) Get(ctx context.Context, id string) (*service.Brand, error) {
	var brand service.Brand
	err := r.db.QueryRowContext(ctx,
		"SELECT id, name, status_id, created_at FROM brands WHERE id = $1",
		id,
	).Scan(&brand.Id, &brand.Name, &brand.StatusId, &brand.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &brand, nil
}

func (r *brandRepo) Update(ctx context.Context, brand *service.Brand) error {
	_, err := r.db.ExecContext(ctx,
		"UPDATE brands SET name = $1, status_id = $2 WHERE id = $3",
		brand.Name, brand.StatusId, brand.Id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *brandRepo) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx,
		"DELETE FROM brands WHERE id = $1",
		id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *brandRepo) GetAll(ctx context.Context, page, limit int64) (*service.BrandResult, error) {
	query := `
			SELECT id, name, status_id, created_at
			FROM brands
			ORDER BY name ASC
			LIMIT $1
	`

	rows, err := r.db.QueryContext(ctx, query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	brands := []*service.Brand{}
	for rows.Next() {
		brand := &service.Brand{}
		err := rows.Scan(
			&brand.Id,
			&brand.Name,
			&brand.StatusId,
			&brand.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		brands = append(brands, brand)
	}

	result := &service.BrandResult{
		Brands: brands,
		Total:  len(brands),
	}

	return result, nil
}
