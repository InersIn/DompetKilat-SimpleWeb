package invoice

import (
	modelsInvoice "DompetKilat-SimpleWeb/models/Finance"
	"context"
	"database/sql"
	"fmt"
)

type invoiceRepositoryImpl struct {
	DB *sql.DB
}

func InvoiceRepository(db *sql.DB) Repository {
	return &invoiceRepositoryImpl{DB: db}
}

func (respository *invoiceRepositoryImpl) CreateInvoice(ctx context.Context, invoice modelsInvoice.Invoce, user_id int) error {
	tx, err := respository.DB.Begin()

	script := "INSERT INTO Invoice(name, amount, tenor, grade, rate, type, user_id) VALUE(?, ?, ?, ?, ?, ?, ?);"

	result, err := tx.ExecContext(ctx, script, invoice.Name, invoice.Amount, invoice.Tenor, invoice.Grade, invoice.Rate, invoice.Type, user_id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	fmt.Println(result.LastInsertId())

	return err
}

func (repository *invoiceRepositoryImpl) GetAllInvoice(ctx context.Context, user_id int, types string) (invoices []modelsInvoice.Invoce, err error) {
	script := "SELECT name, amount, tenor, grade, rate, type FROM Invoice WHERE user_id=? AND type=?;"
	result, err := repository.DB.QueryContext(ctx, script, user_id, types)

	for result.Next() {
		var invoice modelsInvoice.Invoce
		err = result.Scan(&invoice.Name, &invoice.Amount, &invoice.Tenor, &invoice.Grade, &invoice.Rate, &invoice.Type)
		if err != nil {
			fmt.Println(err)
			return
		}
		invoices = append(invoices, invoice)
	}
	return
}
