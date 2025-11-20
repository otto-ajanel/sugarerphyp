package service

import (
    "fmt"

    "github.com/otto-ajanel/backgo_tpdp_np/internal/infra/db"
    "github.com/otto-ajanel/backgo_tpdp_np/internal/model"
    "github.com/otto-ajanel/backgo_tpdp_np/internal/repository"
)

type ProductService struct{
    repo *repository.ProductRepo
}

func NewProductService() *ProductService {
    return &ProductService{repo: repository.NewProductRepo()}
}

// CreateProductRequest representa el body esperado para crear un producto (según PHP)
type CreateProductRequest struct {
    Barcode      string                 `json:"barcode"`
    NameProduct  string                 `json:"nameProduct"`
    CategoryId   map[string]interface{} `json:"categoryId"`
    TypeProduct  int                    `json:"typeproduct"`
    BillingPolicy map[string]interface{} `json:"billingPolicy"`
    SalePrice    float64                `json:"salePrice"`
    TaxSale      float64                `json:"taxsale"`
    CostPrice    float64                `json:"costPrice"`
    TaxPurchase  float64                `json:"taxpurchase"`
    EnableFor    map[string]bool        `json:"enableFor"`
}

func (s *ProductService) CreateProduct(req CreateProductRequest) (map[string]interface{}, error) {
    gdb, err := db.Get()
    if err != nil {
        return nil, fmt.Errorf("db connect error: %w", err)
    }

    tx := gdb.Begin()
    if tx.Error != nil {
        return nil, tx.Error
    }

    // Construir product
    categoryCode := 0
    if v, ok := req.CategoryId["code"]; ok {
        switch t := v.(type) {
        case float64:
            categoryCode = int(t)
        case int:
            categoryCode = t
        }
    }
    billingCode := 0
    if v, ok := req.BillingPolicy["code"]; ok {
        switch t := v.(type) {
        case float64:
            billingCode = int(t)
        case int:
            billingCode = t
        }
    }

    prod := &model.Product{
        ProductSKU: req.Barcode,
        ProductDes: req.NameProduct,
        StatusID:   1,
        CategoryID: categoryCode,
        TypeProductID: req.TypeProduct,
        BillingPolicyID: billingCode,
    }

    if err := s.repo.CreateProduct(tx, prod); err != nil {
        tx.Rollback()
        return nil, err
    }

    // Crear ProductsPrice
    pp := &model.ProductsPrice{
        Price: req.SalePrice,
        TaxPrice: req.TaxSale,
        ProductID: prod.ProductID,
        Purchase: req.CostPrice,
        PurchaseTax: req.TaxPurchase,
    }
    if err := s.repo.CreateProductsPrice(tx, pp); err != nil {
        tx.Rollback()
        return nil, err
    }

    // Module aviability según enableFor
    if req.EnableFor["sale"] {
        ma := &model.ModuleAviability{ModuleID: 12, ProductID: prod.ProductID}
        if err := s.repo.CreateModuleAviability(tx, ma); err != nil {
            tx.Rollback()
            return nil, err
        }
    }
    if req.EnableFor["pos"] {
        ma := &model.ModuleAviability{ModuleID: 26, ProductID: prod.ProductID}
        if err := s.repo.CreateModuleAviability(tx, ma); err != nil {
            tx.Rollback()
            return nil, err
        }
    }
    if req.EnableFor["purchase"] {
        ma := &model.ModuleAviability{ModuleID: 5, ProductID: prod.ProductID}
        if err := s.repo.CreateModuleAviability(tx, ma); err != nil {
            tx.Rollback()
            return nil, err
        }
    }

    if err := tx.Commit().Error; err != nil {
        tx.Rollback()
        return nil, err
    }

    // Devolver producto creado como map
    resp := map[string]interface{}{
        "product": prod,
        "price": pp,
    }
    return resp, nil
}

func (s *ProductService) GetAllProducts() ([]model.Product, error) {
    gdb, err := db.Get()
    if err != nil {
        return nil, fmt.Errorf("db connect error: %w", err)
    }
    return s.repo.GetAllProducts(gdb)
}
