package service

import "context"

func (s *service) CreateProduct(ctx context.Context, product *Product) (*Product, error) {
	product, err := s.productRepo.Create(ctx, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *service) GetProduct(ctx context.Context, productID string) (*Product, error) {
	product, err := s.productRepo.Get(ctx, productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *service) GetProducts(ctx context.Context, filterParams *FilterProducts) (*ProductResult, error) {
	products, err := s.productRepo.GetAll(ctx, filterParams)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *service) UpdateProduct(ctx context.Context, product *Product) error {
	err := s.productRepo.Update(ctx, product)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteProduct(ctx context.Context, id string) error {
	err := s.productRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
