package service

import (
	"fmt"

	"sugarerpgo/internal/infra/db"
	"sugarerpgo/internal/model"
	"sugarerpgo/internal/repository"
)

type StoreService struct {
	repo *repository.StoreRepo
}

func NewStoreService() *StoreService {

	return &StoreService{repo: repository.NewStoreRepo()}

}

func (s *StoreService) GetAllStores() ([]model.Store, error) {
	gdb, err := db.Get()
	if err != nil {
		return nil, fmt.Errorf("db connect error: %w", err)
	}
	repo := repository.NewStoreRepo()
	return repo.GetAll(gdb)
}
