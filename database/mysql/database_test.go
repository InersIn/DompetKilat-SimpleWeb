package mysql

import (
	"context"
	"fmt"
	"testing"
)

type User struct {
	Id    string `json:"id"`
	Nama  string `json:"nama"`
	Email string `json:"email"`
}

func TestDatabases(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	users, err := db.QueryContext(ctx, "SELECT * FROM User")
	if err != nil {
		fmt.Println(err)
	}

	for users.Next() {
		var user User
		err := users.Scan(&user.Id, &user.Nama, &user.Email)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(user)
	}
}
