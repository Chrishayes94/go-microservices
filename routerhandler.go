package main

import (
	"net/http"
)

type BaseHandler interface {
	Book(w http.ResponseWriter, r *http.Request)
	Label(w http.ResponseWriter, r *http.Request)
	Cancel(w http.ResponseWriter, r *http.Request)
}
