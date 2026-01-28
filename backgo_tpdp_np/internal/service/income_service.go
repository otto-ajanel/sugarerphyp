package service

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/otto-ajanel/backgo_tpdp_np/internal/infra/db"
	"github.com/otto-ajanel/backgo_tpdp_np/internal/model"
	"github.com/otto-ajanel/backgo_tpdp_np/internal/repository"
)

type IncomeService struct {
	repo *repository.IncomeRepo
}

func NewIncomeService() *IncomeService {
	return &IncomeService{repo: repository.NewIncomeRepo()}
}

func (s *IncomeService) GetAllIncomes() ([]model.ResultIncome, error) {
	gdb, err := db.Get()
	if err != nil {
		return nil, fmt.Errorf("db connect error", err)

	}
	return s.repo.GetAllIncomes(gdb)
}
func (s *IncomeService) CreateIncome(req map[string]interface{}, c fiber.Ctx) (map[string]interface{}, error) {
	fmt.Print("SERvice income create")
	gdb, err := db.Get()
	if err != nil {
		return nil, fmt.Errorf("db connect error: %w", err)
	}
	tx := gdb.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	income := &model.Income{}
	if v, ok := req["document"].(string); ok {
		income.IncomeDoc = v
	}
	if v, ok := req["dateing"].(string); ok {
		income.IncomeDateIn = v
	}

	userData := c.Locals("userData")
	dataMapUser, _ := userData.(map[string]interface{})

	if idUser, err := dataMapUser["id_user"].(float64); err {
		income.UserID = int(idUser)
	}

	if idStore, ok := req["store"].(map[string]interface{})["code"].(float64); ok {
		income.StoreID = int(idStore)
	}

	if supplier, ok := req["supplier"].(map[string]interface{})["code"].(float64); ok {
		income.SupplierID = int(supplier)
	}

	if err := s.repo.CreateIncome(tx, income); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("create income error: %w", err)
	}
	//

	incomeDet := &model.IncomeDet{}
	detIncomeForm := req["details"].([]interface{})

	for _, v := range detIncomeForm {

		detail := v.(map[string]interface{})

		incomeDet.IncomeID = income.IncomeID
		if productFloat, ok := detail["product"].(float64); ok {
			incomeDet.ProductID = int(productFloat) // Convertir float64 → int
		} else {
			tx.Rollback()
			return nil, fmt.Errorf("product ID missing or invalid")
		}
		incomeDet.IncomeDetCount = detail["quantity"].(float64)
		incomeDet.IncomeDetVal = detail["price"].(float64)

		if err := s.repo.CreateIncomeDet(tx, incomeDet); err != nil {
			fmt.Print("EOrro 2", err)

			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("transaction commit error: %w", err)
	}
	res := map[string]interface{}{
		"id": income.IncomeID,
	}
	return res, nil
}
