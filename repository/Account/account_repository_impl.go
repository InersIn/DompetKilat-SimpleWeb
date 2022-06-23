package account

import (
	models "DompetKilat-SimpleWeb/models/Account"
	"context"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type accountRepositoryImpl struct {
	DB *sql.DB
}

func AccountRepository(db *sql.DB) Repository {
	return &accountRepositoryImpl{DB: db}
}

func (repository *accountRepositoryImpl) Register(ctx context.Context, account models.Account) (models.Account, error) {
	bytePass := []byte(account.Password)
	encrypted, err := bcrypt.GenerateFromPassword(bytePass, 14)
	if err != nil {
		return account, err
	}

	script := "INSERT INTO Users(username, email, password) VALUES (?, ?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, account.Username, account.Email, encrypted)
	if err != nil {
		return account, err
	}

	user_id, err := result.LastInsertId()
	if err != nil {
		return account, err
	}

	account.User_id = int(user_id)

	return account, nil
}

func (repository *accountRepositoryImpl) FindByUsername(ctx context.Context, account models.Account) (models.Account, error) {
	script := "SELECT user_id, password FROM Users WHERE username=?"
	result, err := repository.DB.QueryContext(ctx, script, account.Username)
	if err != nil {
		return account, err
	}

	for result.Next() {
		result.Scan(&account.User_id, &account.Password)
	}

	return account, nil
}
