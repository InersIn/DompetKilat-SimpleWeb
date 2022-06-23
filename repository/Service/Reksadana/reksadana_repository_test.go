package reksadana

import (
	"DompetKilat-SimpleWeb/database/mysql"
	modelsReksadana "DompetKilat-SimpleWeb/models/Finance"
	"context"
	"fmt"
	"testing"
)

func TestGetAllReksana(t *testing.T) {
	fmt.Println("TEST")
	var reksadana []modelsReksadana.Reksadana
	ctx := context.Background()
	reksadanaRepo := ReksadanaRepository(mysql.GetConnection())

	reksadana, err := reksadanaRepo.GetAllReksadana(ctx, 7)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reksadana)
}
