package repository

import (
	"gorm.io/gorm"
)

// PermissionRepo contiene métodos para consultar permisos relacionados con módulos/menus.
type PermissionRepo struct{}

func NewPermissionRepo() *PermissionRepo { return &PermissionRepo{} }

func (r *PermissionRepo) GetPermissionsByUser(db *gorm.DB, userID int) ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	q := db.Table("modules").Select("*").Joins("join menus on menus.id_module = modules.id_module").Joins("join userpermissions on menus.id_menu = userpermissions.id_menu").Where("userpermissions.id_user = ?", userID)

	if err := q.Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}
