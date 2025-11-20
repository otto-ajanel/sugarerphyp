package service

import (
    "fmt"

    "github.com/otto-ajanel/backgo_tpdp_np/internal/infra/db"
    "github.com/otto-ajanel/backgo_tpdp_np/internal/repository"
)

type UserService struct{
    repo *repository.PermissionRepo
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
