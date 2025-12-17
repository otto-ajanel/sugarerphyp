package service

import (
	"fmt"

	"github.com/otto-ajanel/backgo_tpdp_np/internal/infra/db"
	"github.com/otto-ajanel/backgo_tpdp_np/internal/model"
	"github.com/otto-ajanel/backgo_tpdp_np/internal/repository"
)

type AtributeService struct {
	repo *repository.AtributeRepo
}

func NewAtributeService() *AtributeService {
	return &AtributeService{repo: repository.NewAtributeRepo()}
}

type CreateAtributeRequest struct {
	CategoryId       int    `json:"categoryId"`
	AtributeDes      string `json:"atributeDes"`
	AtributeTypedata string `json:"atributeTypedata"`
}

func (s *AtributeService) GetAllAtributes() ([]model.Atribute, error) {
	gdb, err := db.Get()
	if err != nil {
		return nil, fmt.Errorf("db connect error: %w", err)
	}
	return s.repo.GetAll(gdb)
}

func (s *AtributeService) CreateAtribute(req CreateAtributeRequest) (model.Atribute, error) {
	gdb, err := db.Get()
	if err != nil {
		return model.Atribute{}, fmt.Errorf("db connect error: %w", err)
	}
	tx := gdb.Begin()
	if tx.Error != nil {
		return model.Atribute{}, tx.Error
	}
	a := &model.Atribute{
		AtributeDes:      req.AtributeDes,
		AtributeTypedata: req.AtributeTypedata,
	}
	if err := s.repo.Create(tx, a); err != nil {
		tx.Rollback()
		return model.Atribute{}, err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return model.Atribute{}, err
	}
	return *a, nil
}
