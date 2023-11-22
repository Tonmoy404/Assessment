package service

import (
	"context"
	"encoding/json"

	"github.com/Tonmoy404/Assessment/logger"
	"github.com/Tonmoy404/Assessment/util"
)

type service struct {
	brandRepo    BrandRepo
	categoryRepo CategoryRepo
	productRepo  ProductRepo
	supplierRepo SupplierRepo
	cache        Cache
	errRepo      ErrorRepo
}

func NewService(
	brandRepo BrandRepo,
	categoryRepo CategoryRepo,
	productRepo ProductRepo,
	supplierRepo SupplierRepo,
	cache Cache,
	errRepo ErrorRepo,

) Service {
	return &service{
		brandRepo:    brandRepo,
		categoryRepo: categoryRepo,
		productRepo:  productRepo,
		supplierRepo: supplierRepo,
		cache:        cache,
		errRepo:      errRepo,
	}
}

func (s *service) Error(ctx context.Context, code string, description string) *ErrorResponse {
	var errDetail *ErrorDetail

	errString, err := s.cache.Get(code)
	if err != nil {
		logger.Error(ctx, "cannot get from redis", err)
	}
	if len(errString) > 0 {
		err = json.Unmarshal([]byte(errString), &errDetail)
		if err != nil {
			logger.Error(ctx, "cannot unmarshal error detail", err)
		}
	}

	if errDetail != nil && len(errDetail.Code) == 0 {
		return &ErrorResponse{
			Timestamp:   util.GetCurrentTimestamp(),
			Description: description,
			Error:       errDetail,
		}
	}

	errDetail, err = s.errRepo.GetError(ctx, code)
	if err != nil {
		logger.Error(ctx, "cannot get from db", err)
		return &ErrorResponse{
			Timestamp:   util.GetCurrentTimestamp(),
			Description: description,
			Error: &ErrorDetail{
				Code:      code,
				MessageEn: "Not Set",
				MessageBn: "Not Set",
			},
		}
	}

	errResponse := &ErrorResponse{
		Timestamp:   util.GetCurrentTimestamp(),
		Description: description,
		Error:       errDetail,
	}

	return errResponse
}

func (s *service) Response(ctx context.Context, description string, data interface{}) *ResponseData {
	return &ResponseData{
		Timestamp:   util.GetCurrentTimestamp(),
		Description: description,
		Data:        data,
	}
}
