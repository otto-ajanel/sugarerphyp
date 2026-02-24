package repository

import (
	"sugarerpgo/internal/model"

	"gorm.io/gorm"
)

type UserRepo struct{}

func NewUserRepo() *UserRepo { return &UserRepo{} }
func (r *UserRepo) GetAllUsers(db *gorm.DB, page int, perPage int) ([]model.User, int64, error) {

	if page < 1 {
		page = 1
	}
	if perPage <= 0 {
		perPage = 10
	}
	var total int64
	if err := db.Model(&model.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * perPage
	var users []model.User
	if err := db.Model(&model.User{}).Offset(offset).
		Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil

}

func (r *UserRepo) CreateUser(tx *gorm.DB, mu *model.User) error {
	return tx.Create(mu).Error
}
