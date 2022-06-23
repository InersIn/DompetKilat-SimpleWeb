package reksadana

import (
	modelsReksadana "DompetKilat-SimpleWeb/models/Finance"
	"context"
	"database/sql"
	"fmt"
)

type reksadanaRepositoryImpl struct {
	DB *sql.DB
}

func ReksadanaRepository(db *sql.DB) Repository {
	return &reksadanaRepositoryImpl{DB: db}
}

func (repository *reksadanaRepositoryImpl) CreateReksadana(ctx context.Context, reksadana modelsReksadana.Reksadana, user_id int) error {
	tx, err := repository.DB.Begin()

	script := "INSERT INTO Reksadana(`name`, `amount`, `return`, `user_id`) VALUE(?, ?, ?, ?);"

	result, err := tx.ExecContext(ctx, script, reksadana.Name, reksadana.Amount, reksadana.Return, user_id)

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

func (repository *reksadanaRepositoryImpl) GetAllReksadana(ctx context.Context, user_id int) (reksadana []modelsReksadana.Reksadana, err error) {
	script := "SELECT `name`, `amount`, `return`, `user_id` FROM Reksadana WHERE user_id=?;"

	result, err := repository.DB.QueryContext(ctx, script, user_id)
	if err != nil {
		fmt.Println(err)
	}

	for result.Next() {
		fmt.Println(result)
		var RD modelsReksadana.Reksadana
		err = result.Scan(&RD.Name, &RD.Amount, &RD.Return, &RD.User_id)
		if err != nil {
			return
		}
		reksadana = append(reksadana, RD)
	}

	return
}
