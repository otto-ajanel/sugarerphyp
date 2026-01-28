package repository

import (
	"github.com/otto-ajanel/backgo_tpdp_np/internal/model"
	"gorm.io/gorm"
)

type ProductRepo struct{}

func NewProductRepo() *ProductRepo { return &ProductRepo{} }

func (r *ProductRepo) CreateProduct(tx *gorm.DB, p *model.Product) error {
	return tx.Create(p).Error
}

func (r *ProductRepo) CreateProductsPrice(tx *gorm.DB, pp *model.ProductsPrice) error {
	return tx.Create(pp).Error
}

func (r *ProductRepo) CreateModuleAviability(tx *gorm.DB, ma *model.ModuleAviability) error {
	return tx.Create(ma).Error
}

func (r *ProductRepo) GetAllProducts(db *gorm.DB) ([]model.Product, error) {
	var products []model.Product
	if err := db.Preload("Category").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// GetAllProductsPaginated returns products joined with categories and typeproducts with pagination.
// page starts at 1, perPage defaults should be handled by caller.
func (r *ProductRepo) GetAllProductsPaginated(db *gorm.DB, page int, perPage int) ([]model.Product, int64, error) {
	if page < 1 {
		page = 1
	}
	if perPage <= 0 {
		perPage = 10
	}

	var total int64
	if err := db.Model(&model.Product{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * perPage

	var products []model.Product
	// perform joins to include category and typeproduct relations; preload Category to get its description
	if err := db.Model(&model.Product{}).
		Joins("inner join categories c on c.id_category = products.category_id").
		Joins("inner join typeproducts t on t.typeproduct_id = products.typeproduct_id").
		Preload("Category").
		Limit(perPage).
		Offset(offset).
		Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func (r *ProductRepo) CreateProductAttrDet(tx *gorm.DB, pad *model.Product_attridet) error {
	return tx.Create(pad).Error
}
