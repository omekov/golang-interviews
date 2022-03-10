package http

import "net/http"

type Handler struct {
	// service *service
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Init() *http.ServeMux {

	return nil
}
