package service

import (
	"context"
	"time"
)

type Service interface {
	Error(ctx context.Context, internalCode string, description string) *ErrorResponse
	Response(ctx context.Context, description string, data interface{}) *ResponseData
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

type ProductStockRepo interface {
}

type ErrorRepo interface {
	GetError(ctx context.Context, internalCode string) (*ErrorDetail, error)
}
