package invoice

import (
	modelsInvoice "DompetKilat-SimpleWeb/models/Finance"
	"context"
)

type Repository interface {
	CreateInvoice(ctx context.Context, invoice modelsInvoice.Invoce, user_id int) (error)
	GetAllInvoice(ctx context.Context, user_id int, types string) ([]modelsInvoice.Invoce, error)
}
