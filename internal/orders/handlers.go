package orders

import (
	"log"
	"net/http"

	"github.com/APrem-7/GO_ECOM_API/internal/json"
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
	var tempOrder CreateOrderParams
	if err := json.ReadJSON(r, &tempOrder); err != nil {
		log.Println("error")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
