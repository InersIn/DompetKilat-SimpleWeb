package account

import (
	"DompetKilat-SimpleWeb/database/mysql"
	models "DompetKilat-SimpleWeb/models/Account"
	"context"
	"fmt"
	"testing"

	"golang.org/x/crypto/bcrypt"
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

func TestFindByUsername(t *testing.T) {
	login := AccountRepository(mysql.GetConnection())
	ctx := context.Background()
	var account models.Account
	var err error

	loginData := models.Account{
		Username: "Dwii",
		Password: "DwiWardana",
	}

	account, err = login.FindByUsername(ctx, loginData)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(account)

	byteHashed := []byte(account.Password)
	bytePass := []byte(loginData.Password)

	if bcrypt.CompareHashAndPassword(byteHashed, bytePass) != nil {
		t.FailNow()
	}
}
