package repository

import (
	"github.com/otto-ajanel/backgo_tpdp_np/internal/model"
	"gorm.io/gorm"
)

// storeRepo implementa acceso a datos para stores.
type StoreRepo struct{}

func NewStoreRepo() *StoreRepo { return &StoreRepo{} }

func (r *StoreRepo) GetAll(db *gorm.DB) ([]model.Store, error) {
	var stores []model.Store
	if err := db.Find(&stores).Error; err != nil {
		return nil, err
	}
	return stores, nil
}
