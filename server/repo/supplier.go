package repo

import (
	"context"

	"github.com/Tonmoy404/Assessment/service"
	"github.com/jmoiron/sqlx"
)

type SupplierRepo interface {
	service.SupplierRepo
}

type supplierRepo struct {
	db *sqlx.DB
}

func NewSupplierRepo(db *sqlx.DB) SupplierRepo {
	return &supplierRepo{
		db: db,
	}
}

func (r *supplierRepo) Create(ctx context.Context, Supplier *service.Supplier) (*service.Supplier, error) {
	var newSupplier service.Supplier
	err := r.db.QueryRowContext(ctx,
		"INSERT INTO suppliers (name, email, phone, status_id, is_verified_supplier, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *",
		Supplier.Name, Supplier.Email, Supplier.Phone, Supplier.StatusId, Supplier.IsVerifiedSupplier, Supplier.CreatedAt,
	).Scan(&newSupplier.Id, &newSupplier.Name, &newSupplier.Email, &newSupplier.Phone, &newSupplier.StatusId, &newSupplier.IsVerifiedSupplier, &newSupplier.CreatedAt)
	if err != nil {
		return nil, err
	}

	res := &service.Supplier{
		Id:                 newSupplier.Id,
		Name:               newSupplier.Name,
		Email:              newSupplier.Email,
		Phone:              newSupplier.Phone,
		StatusId:           newSupplier.StatusId,
		IsVerifiedSupplier: newSupplier.IsVerifiedSupplier,
		CreatedAt:          newSupplier.CreatedAt,
	}

	return res, nil
}

func (r *supplierRepo) Get(ctx context.Context, id string) (*service.Supplier, error) {
	query := `
			SELECT id, name, email, phone, status_id, is_verified_supplier, created_at
			FROM suppliers
			WHERE id = $1
	`
	supplier := &service.Supplier{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&supplier.Id,
		&supplier.Name,
		&supplier.Email,
		&supplier.Phone,
		&supplier.StatusId,
		&supplier.IsVerifiedSupplier,
		&supplier.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return supplier, nil
}

func (r *supplierRepo) GetAll(ctx context.Context, page, limit int64) (*service.SupplierResult, error) {
	query := `
			SELECT id, name, email, phone, status_id, is_verified_supplier, created_at
			FROM suppliers
			ORDER BY name ASC
			LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, limit, (page-1)*limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	suppliers := []*service.Supplier{}
	for rows.Next() {
		supplier := &service.Supplier{}
		err := rows.Scan(
			&supplier.Id,
			&supplier.Name,
			&supplier.Email,
			&supplier.Phone,
			&supplier.StatusId,
			&supplier.IsVerifiedSupplier,
			&supplier.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		suppliers = append(suppliers, supplier)
	}

	// Create a SupplierResult object with the retrieved suppliers
	result := &service.SupplierResult{
		Suppliers: suppliers,
		Total:     len(suppliers),
	}

	return result, nil
}

func (r *supplierRepo) Update(ctx context.Context, supplier *service.Supplier) error {
	query := `
			UPDATE suppliers
			SET name = $1, email = $2, phone = $3, status_id = $4, is_verified_supplier = $5, created_at = $6
			WHERE id = $7
	`
	_, err := r.db.ExecContext(ctx, query,
		supplier.Id,
		supplier.Name,
		supplier.Email,
		supplier.Phone,
		supplier.StatusId,
		supplier.IsVerifiedSupplier,
		supplier.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *supplierRepo) Delete(ctx context.Context, id string) error {
	query := `
			DELETE FROM suppliers WHERE id = $1
	`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
