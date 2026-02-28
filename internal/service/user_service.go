package service

import (
	"fmt"
	"sugarerpgo/internal/dto/request_dto"
	"sugarerpgo/internal/infra/db"
	"sugarerpgo/internal/model"
	"sugarerpgo/internal/repository"
)

type UserService struct {
	repo *repository.PermissionRepo
}

type UserServCrud struct {
	repo *repository.UserRepo
}

func NewUserService() *UserService {
	return &UserService{repo: repository.NewPermissionRepo()}
}

// GetPermissionsByUser obtiene los permisos (módulos) para un userId
func (s *UserService) GetPermissionsByUser(userID int) ([]map[string]interface{}, error) {
	gdb, err := db.Get()
	if err != nil {
		return nil, fmt.Errorf("db connect error: %w", err)
	}

	perms, err := s.repo.GetPermissionsByUser(gdb, userID)
	if err != nil {
		return nil, err
	}
	return perms, nil
}

func NewUserServCrud() *UserServCrud {
	return &UserServCrud{repo: repository.NewUserRepo()}
}

func (s *UserServCrud) GetUsersPaginate(page int, perPage int) (map[string]interface{}, error) {
	if page < 1 {
		page = 1
	}
	if perPage <= 0 {
		perPage = 10
	}

	gdb, err := db.Get()
	if err != nil {
		return nil, fmt.Errorf("db connect error: %w", err)
	}

	users, total, err := s.repo.GetAllUsers(gdb, page, perPage)
	if err != nil {
		return nil, err
	}
	totalPages := int((total + int64(perPage) - 1) / int64(perPage))

	resp := map[string]interface{}{
		"users":       users,
		"page":        page,
		"per_page":    perPage,
		"total":       total,
		"total_pages": totalPages,
	}
	return resp, nil
}

func (s *UserServCrud) CreateUser(reqUser request_dto.UserReq) (map[string]interface{}, error) {

	gdb, err := db.Get()
	tx := gdb.Begin()
	if err != nil {
		return nil, fmt.Errorf("db connect error: %w", err)
	}

	// Hashear la contraseña antes de guardar
	hashedPassword, err := HashPassword(reqUser.Password)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	newUser := &model.User{
		Name:       reqUser.Name,
		Email:      reqUser.Email,
		Lastname:   reqUser.LastName,
		Password:   hashedPassword,
		IdUsertype: reqUser.TypeUser.Id,
		Active:     true,
		IdTenant:   1,
	}
	if err := s.repo.CreateUser(tx, newUser); err != nil {
		tx.Rollback()
		return nil, err
	}
	newUserStore := &model.UserStore{
		UserId:  newUser.IDUser,
		StoreId: reqUser.Store.Id,
		StateId: 1,
	}
	if err := s.repo.CreateUserStore(tx, newUserStore); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	resp := map[string]interface{}{
		"user": newUser,
	}
	return resp, nil
}

// ChangeUserActive cambia el estado activo de un usuario.
func (s *UserServCrud) ChangeUserActive(userID int, active bool) error {
	gdb, err := db.Get()
	if err != nil {
		return fmt.Errorf("db connect error: %w", err)
	}
	if err := s.repo.UpdateUserActive(gdb, userID, active); err != nil {
		return err
	}
	return nil
}

func (s *UserServCrud) UpdateUser(reqUser request_dto.UserReqUpdate) (map[string]interface{}, error) {

	gdb, err := db.Get()
	tx := gdb.Begin()
	if err != nil {
		return nil, fmt.Errorf("db connect error: %w", err)
	}

	updates := map[string]interface{}{}

	if reqUser.Email != "" {
		updates["email"] = reqUser.Email
	}
	if reqUser.Name != "" {
		updates["name"] = reqUser.Name
	}
	if reqUser.LastName != "" {
		updates["lastname"] = reqUser.LastName
	}
	if reqUser.TypeUser.Id != 0 {
		updates["id_usertype"] = reqUser.TypeUser.Id
	}
	if reqUser.Password != "" {
		hashedPassword, err := HashPassword(reqUser.Password)
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("error hashing password: %w", err)
		}
		updates["password"] = hashedPassword
	}

	if len(updates) > 0 {
		if err := s.repo.UpdateUserFields(tx, reqUser.UserId, updates); err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// si viene store.id, actualizar la relación user_store
	if reqUser.Store.Id != 0 {
		if err := s.repo.UpdateUserStore(tx, reqUser.UserId, reqUser.Store.Id); err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	resp := map[string]interface{}{
		"user_id": reqUser.UserId,
		"updated": updates,
	}
	return resp, nil
}
