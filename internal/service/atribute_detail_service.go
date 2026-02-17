package service

import (
	"fmt"

	"sugarerpgo/internal/infra/db"
	"sugarerpgo/internal/model"
	"sugarerpgo/internal/repository"
)

type AtributeDetailService struct {
	repo *repository.AtributeDetailRepo
}

func NewAtributeDetailService() *AtributeDetailService {
	return &AtributeDetailService{repo: repository.NewAtributeDetailRepo()}
}

type CreateAtributeDetailRequest struct {
	Atribute_id        int    `json:"atributeId"`
	AtriDetDescripcion string `json:"atridet_descripcion"`
}

func (s *AtributeDetailService) GetAllAtributeDetails() ([]model.AtributeDetail, error) {
	gdb, err := db.Get()
	if err != nil {
		return nil, fmt.Errorf("db connect error: %w", err)
	}
	return s.repo.GetAll(gdb)
}

// GetAtributeDetailsByAtributeID devuelve los detalles filtrados por atribute_id.
func (s *AtributeDetailService) GetAtributeDetailsByAtributeID(atributeID int) ([]model.AtributeDetail, error) {
	gdb, err := db.Get()
	if err != nil {
		return nil, fmt.Errorf("db connect error: %w", err)
	}
	return s.repo.GetByAtributeID(gdb, atributeID)
}

func (s *AtributeDetailService) CreateAtributeDetail(req CreateAtributeDetailRequest) (model.AtributeDetail, error) {
	gdb, err := db.Get()
	if err != nil {
		return model.AtributeDetail{}, fmt.Errorf("db connect error: %w", err)
	}
	tx := gdb.Begin()
	if tx.Error != nil {
		return model.AtributeDetail{}, tx.Error
	}
	a := &model.AtributeDetail{
		AtributeID:                req.Atribute_id,
		AtributeDetailDescription: req.AtriDetDescripcion,
	}
	if err := s.repo.Create(tx, a); err != nil {
		tx.Rollback()
		return model.AtributeDetail{}, err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return model.AtributeDetail{}, err
	}
	return *a, nil
}
