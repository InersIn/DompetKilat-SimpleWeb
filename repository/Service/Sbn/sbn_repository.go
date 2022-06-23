package sbn

import (
	modelsSbn "DompetKilat-SimpleWeb/models/Finance"
	"context"
)

type Repository interface {
	CreateSbn(ctx context.Context, sbn modelsSbn.Sbn, user_id int) error
	GetAllSbn(ctx context.Context, user_id int) ([]modelsSbn.Sbn, error)
}
