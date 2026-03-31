package orders

import (
	"context"
	"errors"
	"fmt"

	repo "github.com/APrem-7/GO_ECOM_API/internal/adapters/postgres/sqlc"
	"github.com/jackc/pgx/v5"
)

var (
	ErrorProductNotFound   = errors.New("product not found")
	ErrorProductOutOfStock = errors.New("Product out of stock or quantity not available")
)

type Service interface {
	//define the orders like list order and post order services
	PostOrders(ctx context.Context, tempOrderArg CreateOrderParams) (repo.Order, error)
}

type svc struct {
	repo *repo.Queries
	db   *pgx.Conn
}

func NewService(q *repo.Queries, db *pgx.Conn) Service {
	return &svc{
		repo: q,
		db:   db,
	}
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
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return repo.Order{}, err
	}

	defer tx.Rollback(ctx)
	qtx := s.repo.WithTx(tx)

	//create an order
	order, err := qtx.CreateOrder(ctx, tempOrder.CustomerID)
	if err != nil {
		return repo.Order{}, err
	}

	//looking for the product if exists
	for _, item := range tempOrder.Items {
		product, err := s.repo.GetProductByID(ctx, item.ProductID)
		if err != nil {
			return repo.Order{}, ErrorProductNotFound
		}
		if product.Quantity < item.Quantity {
			return repo.Order{}, ErrorProductOutOfStock
		}
		//error handling of the product

		//creating the order item
		_, err = qtx.CreateOrderItem(ctx, repo.CreateOrderItemParams{
			OrderID:        order.ID,
			Quantity:       item.Quantity,
			PriceInCenters: product.PriceInCenters,
			ProductID:      product.ID,
		})

		if err != nil {
			return repo.Order{}, err
		}

	}
	return order, nil
}
