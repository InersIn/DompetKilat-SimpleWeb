package sbn

import (
	modelsSbn "DompetKilat-SimpleWeb/models/Finance"
	"context"
	"database/sql"
	"fmt"
)

type sbnRepositoryImpl struct {
	DB *sql.DB
}

func SbnRepository(db *sql.DB) Repository {
	return &sbnRepositoryImpl{DB: db}
}

func (repository *sbnRepositoryImpl) CreateSbn(ctx context.Context, sbn modelsSbn.Sbn, user_id int) error {
	tx, err := repository.DB.Begin()

	script := "INSERT INTO SBN(name, amount, tenor, rate, type, user_id) VALUE(?, ?, ?, ?, ?, ?);"

	result, err := tx.ExecContext(ctx, script, sbn.Name, sbn.Amount, sbn.Tenor, sbn.Rate, sbn.Type, user_id)

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

func (repository *sbnRepositoryImpl) GetAllSbn(ctx context.Context, user_id int) (sbn []modelsSbn.Sbn, err error) {
	script := "SELECT name, amount, tenor, rate, type, user_id FROM SBN WHERE user_id=?;"

	result, err := repository.DB.QueryContext(ctx, script, user_id)
	if err != nil {
		fmt.Println(err)
	}

	for result.Next() {
		fmt.Println(result)
		var Sbn modelsSbn.Sbn
		err = result.Scan(&Sbn.Name, &Sbn.Amount, &Sbn.Tenor, &Sbn.Rate, &Sbn.Type, &Sbn.User_id)
		if err != nil {
			return
		}
		sbn = append(sbn, Sbn)
	}

	return
}
