package orders

import (
	"net/http"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}

}

func (h *handler) PostOrder(w http.ResponseWriter, r *http.Request) {
	
}
