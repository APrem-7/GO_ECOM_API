package orders

import (
	"context"
	"fmt"

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

func (s *svc) PostOrders(ctx context.Context, tempOrder CreateOrderParams) (repo.Order, error) {

	//validate the payload
	if tempOrder.CustomerID == 0 {
		return repo.Order{}, fmt.Errorf("The customerID has to be valid")
	}
	if len(tempOrder.Items) == 0 {
		return repo.Order{}, fmt.Errorf("At least one item is required")
	}
	//create an Order
	//look for the product if exists
	//create order item
}
