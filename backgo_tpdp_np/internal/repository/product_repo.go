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
    if err := db.Find(&products).Error; err != nil {
        return nil, err
    }
    return products, nil
}
