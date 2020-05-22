package base

import (
	"github.com/gorilla/mux"
)

func New(serverAddress string, handler BaseHandler) (s Service) {
	s = Service{handler: handler, ServerAddress: serverAddress}
	return s
}

type Service struct {
	ServerAddress		string
	handler				BaseHandler
}

func (s *Service) New() (router *mux.Router) {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/book", s.handler.Book)
	r.HandleFunc("/cancel", s.handler.Cancel)
	r.HandleFunc("/label", s.handler.Label)
	r.HandleFunc("/configuration", s.handler.Configuration)
	return r
}