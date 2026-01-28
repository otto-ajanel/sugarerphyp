package repository

import (
	"fmt"

	"github.com/otto-ajanel/backgo_tpdp_np/internal/model"
	"gorm.io/gorm"
)

type SupplierRepo struct {
}

func NewSupplierRepo() *SupplierRepo {
	return &SupplierRepo{}
}
func (r *SupplierRepo) GetAllSuppliers(db *gorm.DB) ([]model.Supplier, error) {
	var suppliers []model.Supplier
	if err := db.Find(&suppliers).Error; err != nil {
		return nil, fmt.Errorf("error fetching suppliers: %w", err)
	}
	return suppliers, nil
}
