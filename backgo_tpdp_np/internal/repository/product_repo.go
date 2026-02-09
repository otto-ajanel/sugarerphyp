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
func (r *ProductRepo) GetImageProduct(db *gorm.DB, productID int) (string, error) {
	var productPath model.Product_path
	if err := db.Where("product_id = ?", productID).First(&productPath).Error; err != nil {
		return "", err
	}
	return productPath.ProductPath, nil
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

func (r *ProductRepo) GetProductsPaginatedAviable(db *gorm.DB, page int, perPage int) ([]model.ProductAviable, int64, error) {
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

	var products []model.ProductAviable
	// perform joins to include category and typeproduct relations; preload Category to get its description
	if err := db.Table("products p").
		Select(`p.product_id,p.product_sku, p.product_des, c.des_category as category_des, pp.prodprec_price as price, 
		COALESCE(tot.total_available, 0) AS total_available,
		CASE WHEN COALESCE(tot.total_available, 0) > 0 THEN 'INSTOCK' ELSE 'OUTOFSTOCK' END AS inventory_status,

        (
            SELECT json_agg(
                json_build_object(
                    'store_id', store_id,
                    'count_available', count_available,
					'store_name', store_name

                )
            )
            FROM (
                SELECT
                    s.id_store AS store_id,
					s.store_name AS store_name,
                    SUM(id.incomedet_count) AS count_available
                FROM income_det id
                INNER JOIN incomes i  ON i.income_id = id.income_id
                INNER JOIN stores s ON s.id_store = i.store_id
                WHERE id.product_id = p.product_id
                GROUP BY s.id_store
            ) t
        ) AS locations
		`).
		Joins("inner join categories c on c.id_category = p.category_id").
		Joins("inner join typeproducts t on t.typeproduct_id = p.typeproduct_id").
		Joins("inner join products_prices pp on pp.product_id = p.product_id").
		Joins(`
		LEFT JOIN (
			SELECT
				id.product_id,
				SUM(id.incomedet_count) AS total_available
			FROM income_det id
			GROUP BY id.product_id
		) tot ON tot.product_id = p.product_id
		`).
		Where("pp.status_id = 1").
		Limit(perPage).
		Offset(offset).
		Scan(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}
