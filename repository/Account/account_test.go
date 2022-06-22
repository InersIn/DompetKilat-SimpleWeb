package account

import (
	"DompetKilat-SimpleWeb/database/mysql"
	models "DompetKilat-SimpleWeb/models/Account"
	"context"
	"fmt"
	"testing"
)

func TestRegisterAccount(t *testing.T) {
	register := AccountRepository(mysql.GetConnection())

	ctx := context.Background()

	account := models.Account{
		Username: "Dwii",
		Password: "DwiWardana",
		Email:    "sapu2776@gmail.com",
	}

	account, err := register.Register(ctx, account)
	if err != nil {
		fmt.Println(err)
	}
}
