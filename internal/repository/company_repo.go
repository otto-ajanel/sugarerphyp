package repository

import (
	"sugarerpgo/internal/model"

	"gorm.io/gorm"
)

type CompanyRepo struct{}

func NewCompanyRepo() *CompanyRepo { return &CompanyRepo{} }
func (r *CompanyRepo) GetAll(db *gorm.DB) ([]model.Company, error) {
	var companies []model.Company
	if err := db.Find(&companies).Error; err != nil {
		return nil, err
	}
	return companies, nil
}
