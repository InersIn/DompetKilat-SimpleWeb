package finance

import (
	modelsFinane "DompetKilat-SimpleWeb/models/Finance"
	"context"
)

type Repository interface {
	CreateFinance(ctx context.Context, invoice modelsFinane.Finance, user_id int) error
	GetAllFinance(ctx context.Context, user_id int) ([]modelsFinane.Finance, error)
}
