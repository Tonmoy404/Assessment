package service

import (
	"context"
	"time"
)

type Service interface {
	Error(ctx context.Context, internalCode string, description string) *ErrorResponse
	Response(ctx context.Context, description string, data interface{}) *ResponseData

	///Brand Services
	CreateBrand(ctx context.Context, brand *Brand) (*Brand, error)
	GetBrand(ctx context.Context, id string) (*Brand, error)
	GetBrands(ctx context.Context, page, limit int64) (*BrandResult, error)
	UpdateBrand(ctx context.Context, brand *Brand) error
	DeleteBrand(ctx context.Context, id string) error

	///Product Services
	CreateProduct(ctx context.Context, product *Product) (*Product, error)
	GetProduct(ctx context.Context, name string) (*Product, error)
	GetProducts(ctx context.Context, param *FilterProducts) (*ProductResult, error)
	UpdateProduct(ctx context.Context, product *Product) error
	DeleteProduct(ctx context.Context, id string) error

	///Supplier Services
	CreateSupplier(ctx context.Context, newSupplier *Supplier) (*Supplier, error)
	GetSupplier(ctx context.Context, id string) (*Supplier, error)
	UpdateSupplier(ctx context.Context, supplier *Supplier) error
	DeleteSupplier(ctx context.Context, id string) error
	GetSuppliers(ctx context.Context, page, limit int64) (*SupplierResult, error)
}

type Cache interface {
	Set(key string, value string, ttl time.Duration) error
	Get(key string) (string, error)
	Delete(key string) error
	GetTTL(key string) (time.Duration, error)
}

type BrandRepo interface {
	Create(ctx context.Context, brand *Brand) (*Brand, error)
	Get(ctx context.Context, id string) (*Brand, error)
	GetAll(ctx context.Context, page, limit int64) (*BrandResult, error)
	Update(ctx context.Context, brand *Brand) error
	Delete(ctx context.Context, id string) error
}

type CategoryRepo interface {
}

type ProductRepo interface {
	Create(ctx context.Context, product *Product) (*Product, error)
	Get(ctx context.Context, id string) (*Product, error)
	GetAll(ctx context.Context, filterParams *FilterProducts) (*ProductResult, error)
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id string) error
}

type SupplierRepo interface {
	Create(ctx context.Context, supplier *Supplier) (*Supplier, error)
	Get(ctx context.Context, id string) (*Supplier, error)
	GetAll(ctx context.Context, page, limit int64) (*SupplierResult, error)
	Update(ctx context.Context, supplier *Supplier) error
	Delete(ctx context.Context, id string) error
}

type ErrorRepo interface {
	GetError(ctx context.Context, internalCode string) (*ErrorDetail, error)
}
