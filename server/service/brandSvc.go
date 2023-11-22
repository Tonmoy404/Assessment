package service

import "context"

func (s *service) CreateBrand(ctx context.Context, brand *Brand) (*Brand, error) {
	newBrand, err := s.brandRepo.Create(ctx, brand)
	if err != nil {
		return nil, err
	}

	return newBrand, nil
}

func (s *service) GetBrand(ctx context.Context, id string) (*Brand, error) {
	brand, err := s.brandRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return brand, nil
}

func (s *service) GetBrands(ctx context.Context, page, limit int64) (*BrandResult, error) {
	result, err := s.brandRepo.GetAll(ctx, page, limit)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) UpdateBrand(ctx context.Context, brand *Brand) error {
	err := s.brandRepo.Update(ctx, brand)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteBrand(ctx context.Context, id string) error {
	err := s.brandRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
