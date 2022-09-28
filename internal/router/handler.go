package router

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

func (s *Server) GetOrders(w http.ResponseWriter, req *http.Request) {
	orders, err := s.service.GetOrders()
	if err != nil {
		log.Println(err)
		errorRes(w, http.StatusBadRequest, "server suffer")
		return
	}

	if len(orders) == 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]int{})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(orders)
	if err != nil {
		errorRes(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (s *Server) GetOrderByUuid(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["uuid"]

	order, err := s.service.GetOrderByUUID(uuid.MustParse(id))
	if err != nil {
		log.Println(err)
		errorRes(w, http.StatusBadRequest, "pepa")
		return
	}

	if order == nil {
		errorRes(w, http.StatusNoContent, "no content")
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(order)
	if err != nil {
		errorRes(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func errorRes(w http.ResponseWriter, statusCode int, msg string) {
	msgWithErr := errorResponse{Message: msg}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if msg != "" {
		_ = json.NewEncoder(w).Encode(msgWithErr)
	}
}
