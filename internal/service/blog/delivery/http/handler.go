package http

type IHandler interface{}

type Handler struct {
	//usecase here
}

func NewHandler() IHandler {
	return &Handler{}
}
