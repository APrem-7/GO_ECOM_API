package products

import (
	"context"

	repo "github.com/APrem-7/GO_ECOM_API/internal/adapters/postgres/sqlc"
)

type Service interface {
	//define the listProductsService
	ListProduct(ctx context.Context) ([]repo.Product, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(q repo.Querier) Service {
	return &svc{repo: q}
}

func (s *svc) ListProduct(ctx context.Context) ([]repo.Product, error) {
	return s.repo.ListProducts(ctx)

}
