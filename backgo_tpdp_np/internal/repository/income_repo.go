package repository

import (
	"github.com/otto-ajanel/backgo_tpdp_np/internal/model"
	"gorm.io/gorm"
)

type IncomeRepo struct{}

func NewIncomeRepo() *IncomeRepo { return &IncomeRepo{} }

func (r *IncomeRepo) GetAllIncomes(db *gorm.DB) ([]model.ResultIncome, error) {
	var resultIncomes []model.ResultIncome

	if err := db.Table("incomes").
		Select("incomes.income_id, incomes.income_doc, incomes.income_dateing, u.username, s.store_name").
		Joins("inner join users u on u.id_user = incomes.user_id").
		Joins("inner join stores s on s.id_store = incomes.store_id").
		Find(&resultIncomes).Error; err != nil {
		return nil, err
	}
	return resultIncomes, nil
}

func (r *IncomeRepo) CreateIncome(tx *gorm.DB, income *model.Income) error {
	return tx.Create(income).Error
}

func (r *IncomeRepo) CreateIncomeDet(tx *gorm.DB, incomeDet *model.IncomeDet) error {
	return tx.Create(incomeDet).Error
}
