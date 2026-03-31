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
	//calling the service here for creating the order
	createdOrder, err := h.service.PostOrders(r.Context(), tempOrder)
	if err != nil {
		log.Println("error")
		if err == ErrorProductNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return

		}
		if err == ErrorProductOutOfStock {
			http.Error(w, err.Error(), http.StatusInsufficientStorage)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.WriteJSON(w, http.StatusCreated, createdOrder)

}
