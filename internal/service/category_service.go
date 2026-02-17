package service

import (
	"fmt"

	"sugarerpgo/internal/infra/db"
	"sugarerpgo/internal/model"
	"sugarerpgo/internal/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepo
}

func NewCategoryService() *CategoryService {
	return &CategoryService{repo: repository.NewCategoryRepo()}
}

type CreateCategoryRequest struct {
	NewCategory string `json:"newCategory"`
}

func (s *CategoryService) GetAllCategories() ([]model.Category, error) {
	gdb, err := db.Get()
	if err != nil {
		return nil, fmt.Errorf("db connect error: %w", err)
	}
	return s.repo.GetAll(gdb)
}

func (s *CategoryService) CreateCategory(req CreateCategoryRequest) (model.Category, error) {
	gdb, err := db.Get()
	if err != nil {
		return model.Category{}, fmt.Errorf("db connect error: %w", err)
	}

	tx := gdb.Begin()
	if tx.Error != nil {
		return model.Category{}, tx.Error
	}

	cat := &model.Category{Des: req.NewCategory}
	if err := s.repo.Create(tx, cat); err != nil {
		tx.Rollback()
		return model.Category{}, err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return model.Category{}, err
	}
	return *cat, nil
}
