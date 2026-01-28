package service

import (
	"fmt"

	"github.com/otto-ajanel/backgo_tpdp_np/internal/infra/db"
	"github.com/otto-ajanel/backgo_tpdp_np/internal/model"
	"github.com/otto-ajanel/backgo_tpdp_np/internal/repository"
)

type CompanyService struct {
	repo *repository.CompanyRepo
}

func NewCompanyService() *CompanyService {
	return &CompanyService{repo: repository.NewCompanyRepo()}
}

func (s *CompanyService) GetAllCompanies() ([]model.Company, error) {
	gdb, err := db.Get()
	if err != nil {
		return nil, fmt.Errorf("db connect error: %w", err)
	}
	return s.repo.GetAll(gdb)
}
