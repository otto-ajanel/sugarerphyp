package repository

import (
	"sugarerpgo/internal/model"

	"gorm.io/gorm"
)

type ProductPathRepo struct{}

func NewProductPathRepo() *ProductPathRepo { return &ProductPathRepo{} }
func (r *ProductPathRepo) Create(tx *gorm.DB, pp *model.Product_path) error {
	return tx.Create(pp).Error
}

func (r *ProductPathRepo) SaveProductPath(tx *gorm.DB, pp *model.Product_path) error {
	return tx.Save(pp).Error
}
