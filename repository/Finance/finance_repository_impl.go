package finance

import (
	modelsFinane "DompetKilat-SimpleWeb/models/Finance"
	"context"
	"database/sql"
	"fmt"
)

type financeRepositoryImpl struct {
	DB *sql.DB
}

func FinanceRepository(db *sql.DB) Repository {
	return &financeRepositoryImpl{DB: db}
}

func (repository *financeRepositoryImpl) CreateFinance(ctx context.Context, finance modelsFinane.Finance, user_id int) (error) {
	tx, err := repository.DB.Begin()

	script := "INSERT INTO Finance(name, count, sub, user_id) VALUES(?, ?, ?, ?);"

	var sub interface{}

	if finance.Sub == "" {
		sub = nil
	}

	result, err := tx.ExecContext(ctx, script, finance.Name, finance.Count, sub, user_id)
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
