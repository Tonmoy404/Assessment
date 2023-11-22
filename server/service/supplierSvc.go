package service

import "context"

func (s *service) CreateSupplier(ctx context.Context, supplier *Supplier) (*Supplier, error) {
	newSupplier, err := s.supplierRepo.Create(ctx, supplier)
	if err != nil {
		return nil, err
	}

	return newSupplier, nil
}

func (s *service) GetSupplier(ctx context.Context, id string) (*Supplier, error) {
	supplier, err := s.supplierRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return supplier, nil
}

func (s *service) GetSuppliers(ctx context.Context, page, limit int64) (*SupplierResult, error) {
	result, err := s.supplierRepo.GetAll(ctx, page, limit)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) UpdateSupplier(ctx context.Context, spplr *Supplier) error {
	err := s.supplierRepo.Update(ctx, spplr)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteSupplier(ctx context.Context, id string) error {
	err := s.supplierRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
