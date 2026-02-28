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

// UpdateUserActive modifica el campo active de un usuario existente.
func (r *UserRepo) UpdateUserActive(db *gorm.DB, userID int, active bool) error {
	return db.Model(&model.User{}).
		Where("id_user = ?", userID).
		Update("active", active).
		Error
}

func (r *UserRepo) CreateUserStore(db *gorm.DB, mus *model.UserStore) error {
	return db.Create(mus).
		Error
}

// UpdateUserFields actualiza campos del usuario usando la transacción proporcionada.
func (r *UserRepo) UpdateUserFields(tx *gorm.DB, userID int, updates map[string]interface{}) error {
	return tx.Model(&model.User{}).
		Where("id_user = ?", userID).
		Updates(updates).
		Error
}

// UpdateUserStore actualiza la relación user_store; si no existe, la crea.
func (r *UserRepo) UpdateUserStore(tx *gorm.DB, userID int, storeID int) error {
	res := tx.Model(&model.UserStore{}).
		Where("user_id = ?", userID).
		Updates(map[string]interface{}{"store_id": storeID})

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		// no existía, crear nueva relación
		us := &model.UserStore{UserId: userID, StoreId: storeID, StateId: 1}
		if err := tx.Create(us).Error; err != nil {
			return err
		}
	}
	return nil
}
