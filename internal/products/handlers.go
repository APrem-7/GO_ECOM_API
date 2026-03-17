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