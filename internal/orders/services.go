package orders

import (
	"context"

	repo "github.com/APrem-7/GO_ECOM_API/internal/adapters/postgres/sqlc"
)

type Service interface {
	//define the orders like list order and post order services
	PostOrders(ctx context.Context, tempOrderArg CreateOrderParams) (repo.Order, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(q repo.Querier) Service {
	return &svc{repo: q}
}
