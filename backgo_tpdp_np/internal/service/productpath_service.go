package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

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

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// If an image file was provided, save it to the `imagenes` dir and update the record

	if req.ImageFile != nil {
		fmt.Println("Saving image file...")
		// ensure directory exists
		dir := "images"
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("create images dir: %w", err)
		}

		// derive extension from uploaded filename
		ext := filepath.Ext(req.ImageFile.Filename)
		if ext == "" {
			ext = ".img"
		}

		// build filename using product id and productpath id
		filename := fmt.Sprintf("product_%d_%d%s", productPath.ProductID, productPath.Product_pathID, ext)
		dstPath := filepath.Join(dir, filename)

		src, err := req.ImageFile.Open()
		if err != nil {
			return nil, fmt.Errorf("open uploaded file: %w", err)
		}
		defer src.Close()

		dst, err := os.Create(dstPath)
		if err != nil {
			return nil, fmt.Errorf("create destination file: %w", err)
		}
		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			return nil, fmt.Errorf("save file: %w", err)
		}

		// update DB record with saved path
		productPath.ProductPath = dstPath
		if err := s.repo.SaveProductPath(gdb, productPath); err != nil {
			return nil, fmt.Errorf("update product path in db: %w", err)
		}
	}

	resp := map[string]interface{}{
		"message": "Imagen subida correctamente",
	}
	return resp, nil
}
