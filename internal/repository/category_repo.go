package repository

import (
	"sugarerpgo/internal/model"

	"gorm.io/gorm"
)

// CategoryRepo implementa acceso a datos para categories.
type CategoryRepo struct{}

func NewCategoryRepo() *CategoryRepo { return &CategoryRepo{} }

func (r *CategoryRepo) GetAll(db *gorm.DB) ([]model.Category, error) {
	var cats []model.Category
	if err := db.Find(&cats).Error; err != nil {
		return nil, err
	}
	return cats, nil
}

func (r *CategoryRepo) Create(tx *gorm.DB, c *model.Category) error {
	return tx.Create(c).Error
}
