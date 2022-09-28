package router

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"testgo/internal/cache"
	"testgo/internal/service"
)

type Server struct {
	service *service.Service
	cache   *cache.Cache
}

func New(service *service.Service) *Server {
	return &Server{
		service: service,
	}
}

func (s *Server) Start() error {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/api/v1/orders", s.GetOrders).Methods("GET")
	r.HandleFunc("/api/v1/order/{uuid}", s.GetOrderByUuid).Methods("GET")

	return http.ListenAndServe(":8080", handlers.CORS()(r))
}
