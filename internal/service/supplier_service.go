package service

import (
	"fmt"

	"sugarerpgo/internal/infra/db"
	"sugarerpgo/internal/model"
	"sugarerpgo/internal/repository"
)

type SupplierService struct {
	repo *repository.SupplierRepo
}

func NewSupplierService() *SupplierService {
	return &SupplierService{repo: repository.NewSupplierRepo()}
}
func (s *SupplierService) GetAllSuppliers() ([]model.Supplier, error) {
	gdb, err := db.Get()
	if err != nil {
		return nil, fmt.Errorf("db connect error", err)
	}
	return s.repo.GetAllSuppliers(gdb)
}
