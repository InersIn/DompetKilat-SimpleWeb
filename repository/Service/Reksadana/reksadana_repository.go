package reksadana

import (
	modelsReksadana "DompetKilat-SimpleWeb/models/Finance"
	"context"
)

type Repository interface {
	CreateReksadana(ctx context.Context, reksadana modelsReksadana.Reksadana, user_id int) (error)
	GetAllReksadana(ctx context.Context, user_id int) ([]modelsReksadana.Reksadana, error)
}
