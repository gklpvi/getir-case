package server

import (
	"getir-case/service"
	"net/http"
)

type Server struct {
	service *service.Service
}

func New(service *service.Service) *Server {
	return &Server{
		service: service,
	}
}

func (s *Server) Start() {
	http.HandleFunc("/records", s.GetRecordsFromDB)
	http.HandleFunc("/all-records", s.GetAllRecordsFromDB)

	http.HandleFunc("/in-memory", s.InMemoryHandler)
	http.HandleFunc("/all-in-memory", s.GetAllRecordsFromIM)

	http.ListenAndServe(":8080", nil)
}
