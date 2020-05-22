package main

import (
	"github.com/gorilla/mux"
	"go-microservices/handlers"
)

func New(serverAddress string, handler handlers.BaseHandler) (s Service) {
	s = Service{handler: handler, ServerAddress: serverAddress}
	return s
}

type Service struct {
	ServerAddress		string
	handler				handlers.BaseHandler
}

func (s *Service) New() (router *mux.Router) {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/book", s.handler.Book)
	r.HandleFunc("/cancel", s.handler.Cancel)
	r.HandleFunc("/label", s.handler.Label)
	return r
}