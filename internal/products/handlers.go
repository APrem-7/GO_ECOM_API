package products

import (
	"net/http"
)

type handler struct {
	service Service
}

// create a handler in using the Servicce
func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}

}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	//Call the service to List all the products available
	//return JSON in an http Resposne

}
