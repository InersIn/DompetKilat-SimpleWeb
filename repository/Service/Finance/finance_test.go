package finance

import (
	"DompetKilat-SimpleWeb/database/mysql"
	modelsFinance "DompetKilat-SimpleWeb/models/Finance"
	"context"
	"fmt"
	"testing"
)

func TestCreateFinances(t *testing.T) {
	financeRepo := FinanceRepository(mysql.GetConnection())
	ctx := context.Background()

	var finance modelsFinance.Finance
	finance.Name = "Invoice Incoming"
	finance.Count = 35
	finance.Sub = ""

	err := financeRepo.CreateFinance(ctx, finance, 7)
	if err != nil {
		fmt.Println(err)
	}
}

func TestGetAllFinances(t *testing.T) {
	financeRepo := FinanceRepository(mysql.GetConnection())
	ctx := context.Background()

	var finances []modelsFinance.Finance
	finances, err := financeRepo.GetAllFinance(ctx, 7)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(finances)
}
