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
	http.HandleFunc("/records", s.service.GetRecordsFromDB)
	http.HandleFunc("/all-records", s.service.GetAllRecordsFromDB)

	http.HandleFunc("/in-memory", s.service.InMemoryHandler)
	http.HandleFunc("/all-in-memory", s.service.GetAllRecordsFromIM)

	http.ListenAndServe(":8080", nil)
}
