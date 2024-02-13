package server

import (
	"getir-case/service"
	"net/http"
)

type Server struct {
	service *service.Service
}

func NewServer(service *service.Service) *Server {
	return &Server{
		service: service,
	}
}

func (s *Server) Start() {
	http.HandleFunc("/from-db", s.service.GetRecordsFromDB)
	http.HandleFunc("/all-from-db", s.service.GetRecordsFromDB)

	http.ListenAndServe(":8080", nil)
}
