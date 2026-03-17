package products

type handler struct{
	servcie Service,
}


//create a handler in using the Servicce
func NewHandler ( service Service) *handler{
	return &handler{
		servcie: service,
	}

}

func (h * handler) ListProducts(w http.ResponseWriter, r *http.Request){
	//Call the service to List all the products available
	//return JSON in an http Resposne	
	
}