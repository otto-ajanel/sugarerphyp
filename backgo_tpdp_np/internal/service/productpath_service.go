package service

import (
	"fmt"
	"mime/multipart"

	"github.com/otto-ajanel/backgo_tpdp_np/internal/infra/db"
	"github.com/otto-ajanel/backgo_tpdp_np/internal/model"
	"github.com/otto-ajanel/backgo_tpdp_np/internal/repository"
)

type UploadProductImageRequest struct {
	ProductID int                   `form:"productId"`
	ImagePath string                `form:"imagePath"`
	ImageFile *multipart.FileHeader `form:"imageFile"`
}

type ProductPathService struct {
	repo *repository.ProductPathRepo
}

func NewProductPathService() *ProductPathService {
	return &ProductPathService{repo: repository.NewProductPathRepo()}
}
func (s *ProductPathService) UploadProductImage(req UploadProductImageRequest) (map[string]interface{}, error) {
	gdb, err := db.Get()

	if err != nil {
		return nil, fmt.Errorf("db connect error: %w", err)
	}

	tx := gdb.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	productPath := &model.Product_path{
		ProductID:   req.ProductID,
		ProductPath: req.ImagePath,
	}

	if err := s.repo.SaveProductPath(tx, productPath); err != nil {
		return nil, err

	}
	resp := map[string]interface{}{
		"message": "Imagen subida correctamente",
	}
	return resp, nil
}
