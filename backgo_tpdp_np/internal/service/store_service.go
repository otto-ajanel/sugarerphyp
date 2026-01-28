package service

import (
	"fmt"

	"github.com/otto-ajanel/backgo_tpdp_np/internal/infra/db"
	"github.com/otto-ajanel/backgo_tpdp_np/internal/model"
	"github.com/otto-ajanel/backgo_tpdp_np/internal/repository"
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
