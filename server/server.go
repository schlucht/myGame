package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	Port    string
	Message string
	Static  string
	Routes  []Route
}

func (m *Server) Init() {

	http.Handle(m.Static, http.StripPrefix(m.Static, m.staticHandler()))
	for _, route := range m.Routes {
		http.HandleFunc(route.RouterName, route.Handler)
	}
	fmt.Printf("Server run on localhost:%s\n", m.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", m.Port), nil)
	if err != nil {
		panic("Error: " + err.Error())
	}
}

func (m *Server) staticHandler() http.Handler {
	return http.FileServer(http.Dir(fmt.Sprintf("./%s", m.Static)))
}
