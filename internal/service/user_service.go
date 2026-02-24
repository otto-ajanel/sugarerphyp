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

	newUser := &model.User{
		Name:     reqUser.Name,
		Email:    reqUser.Email,
		Lastname: reqUser.LastName,
		Password: reqUser.Password,
	}
	if err := s.repo.CreateUser(tx, newUser); err != nil {
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
