package repo

import (
	"context"
	"time"

	"github.com/Tonmoy404/Assessment/service"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
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

func (r *productRepo) Create(ctx context.Context, product *service.Product) (*service.Product, error) {

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	const insertProductQuery = `
        INSERT INTO products (id, name, description, specification, brand_id, category_id, supplier_id, unit_price, discount_price, tags, status_id, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	_, err = tx.ExecContext(ctx,
		insertProductQuery,
		product.Id,
		product.Name,
		product.Description,
		product.Specification,
		product.BrandId,
		product.CategoryId,
		product.SupplierId,
		product.UnitPrice,
		product.DiscountPrice,
		pq.Array(product.Tags),
		product.StatusId,
		product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	const insertProductStockQuery = `
			INSERT INTO product_stocks (product_id, stock_quantity, updated_at)
			VALUES ($1, $2, $3)
	`
	_, err = tx.ExecContext(ctx,
		insertProductStockQuery,
		product.Id,
		product.Stock,
		time.Now(),
	)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *productRepo) Get(ctx context.Context, id string) (*service.Product, error) {

	const getProductQuery = `
			SELECT id, name, description, specification, brand_id, category_id, supplier_id, unit_price, discount_price, tags, status_id, stock, created_at
			FROM products
			WHERE id = $1
	`
	row := r.db.QueryRowContext(ctx, getProductQuery, id)

	product := &service.Product{}
	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Description,
		&product.Specification,
		&product.BrandId,
		&product.CategoryId,
		&product.SupplierId,
		&product.UnitPrice,
		&product.DiscountPrice,
		pq.Array(&product.Tags),
		&product.StatusId,
		&product.Stock,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *productRepo) Delete(ctx context.Context, id string) error {

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	const deleteProductQuery = `
			DELETE FROM products
			WHERE id = $1
	`
	_, err = tx.ExecContext(ctx, deleteProductQuery, id)
	if err != nil {
		return err
	}

	const deleteProductStockQuery = `
			DELETE FROM product_stocks
			WHERE product_id = $1
	`
	_, err = tx.ExecContext(ctx, deleteProductStockQuery, id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *productRepo) GetAll(ctx context.Context, filterParams *service.FilterProducts) (*service.ProductResult, error) {
	result := &service.ProductResult{}

	countQuery, countArgs := buildCountQuery(filterParams)
	row := r.db.QueryRowContext(ctx, countQuery, countArgs...)
	err := row.Scan(&result.Total)
	if err != nil {
		return nil, err
	}

	selectQuery, selectArgs := buildSelectQuery(filterParams)
	rows, err := r.db.QueryContext(ctx, selectQuery, selectArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		product := &service.Product{}
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Specification,
			&product.BrandId,
			&product.CategoryId,
			&product.SupplierId,
			&product.UnitPrice,
			&product.DiscountPrice,
			pq.Array(&product.Tags),
			&product.StatusId,
			&product.Stock,
			&product.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		result.Products = append(result.Products, *product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *productRepo) Update(ctx context.Context, product *service.Product) error {
	query := `
			UPDATE products
			SET name = $1, description = $2, specification = $3, brand_id = $4, category_id = $5, supplier_id = $6,
					unit_price = $7, discount_price = $8, tags = $9, status_id = $10, stock = $11
			WHERE id = $12
	`
	_, err := r.db.ExecContext(ctx, query,
		product.Name,
		product.Description,
		product.Specification,
		product.BrandId,
		product.CategoryId,
		product.SupplierId,
		product.UnitPrice,
		product.DiscountPrice,
		pq.Array(product.Tags),
		product.StatusId,
		product.Stock,
		product.Id,
	)
	if err != nil {
		return err
	}

	return nil
}

func buildCountQuery(filterParams *service.FilterProducts) (string, []interface{}) {
	// Initialize the query and arguments
	query := "SELECT COUNT(*) FROM products"
	args := []interface{}{}

	// Add WHERE clauses for each filter parameter that is set
	if filterParams != nil {
		if filterParams.BrandId != "" {
			query += " WHERE brand_id = $1"
			args = append(args, filterParams.BrandId)
		}
		if filterParams.CategoryId != "" {
			if len(args) > 0 {
				query += " AND"
			} else {
				query += " WHERE"
			}
			query += " category_id = $2"
			args = append(args, filterParams.CategoryId)
		}
		if filterParams.SupplierId != "" {
			if len(args) > 0 {
				query += " AND"
			} else {
				query += " WHERE"
			}
			query += " supplier_id = $3"
			args = append(args, filterParams.SupplierId)
		}
		if len(filterParams.Tags) > 0 {
			if len(args) > 0 {
				query += " AND"
			} else {
				query += " WHERE"
			}
			query += " tags && $4"
			args = append(args, pq.Array(filterParams.Tags))
		}
	}

	return query, args
}

func buildSelectQuery(filterParams *service.FilterProducts) (string, []interface{}) {
	query := "SELECT id, name, description, specification, brand_id, category_id, supplier_id, unit_price, discount_price, tags, status_id, stock, created_at FROM products"
	args := []interface{}{}

	if filterParams != nil {
		if filterParams.BrandId != "" {
			query += " WHERE brand_id = $1"
			args = append(args, filterParams.BrandId)
		}
		if filterParams.CategoryId != "" {
			if len(args) > 0 {
				query += " AND"
			} else {
				query += " WHERE"
			}
			query += " category_id = $2"
			args = append(args, filterParams.CategoryId)
		}
		if filterParams.SupplierId != "" {
			if len(args) > 0 {
				query += " AND"
			} else {
				query += " WHERE"
			}
			query += " supplier_id = $3"
			args = append(args, filterParams.SupplierId)
		}
		if len(filterParams.Tags) > 0 {
			if len(args) > 0 {
				query += " AND"
			} else {
				query += " WHERE"
			}
			query += " tags && $4"
			args = append(args, pq.Array(filterParams.Tags))
		}
	}
	query += " ORDER BY created_at DESC LIMIT $5"
	args = append(args, filterParams.Limit)

	return query, args
}
