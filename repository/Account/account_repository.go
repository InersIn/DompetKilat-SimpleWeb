package account

import (
	models "DompetKilat-SimpleWeb/models/Account"
	"context"
)

type Repository interface {
	Register(ctx context.Context, account models.Account) (models.Account, error)
	FindByUsername(ctx context.Context, username string) (models.Account, error)
}
