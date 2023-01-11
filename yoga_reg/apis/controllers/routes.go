package controllers

import (
	//"net/http"

	"github.com/ha-wk/yoga_reg/apis/middlewares"
)

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//payments
	s.Router.HandleFunc("/pay/{id}", middlewares.SetMiddlewareJSON(s.App)).Methods("PUT")
}
