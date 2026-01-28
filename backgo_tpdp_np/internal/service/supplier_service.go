package service

import (
	"fmt"

	"github.com/otto-ajanel/backgo_tpdp_np/internal/infra/db"
	"github.com/otto-ajanel/backgo_tpdp_np/internal/model"
	"github.com/otto-ajanel/backgo_tpdp_np/internal/repository"
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
