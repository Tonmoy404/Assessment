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
}

type Cache interface {
	Set(key string, value string, ttl time.Duration) error
	Get(key string) (string, error)
	Delete(key string) error
	GetTTL(key string) (time.Duration, error)
}

type BrandRepo interface {
}

type CategoryRepo interface {
}

type ProductRepo interface {
}

type SupplierRepo interface {
}

type ErrorRepo interface {
	GetError(ctx context.Context, internalCode string) (*ErrorDetail, error)
}
