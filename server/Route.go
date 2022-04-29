package server

import "net/http"

type Route struct {
	RouterName string
	Handler    func(http.ResponseWriter, *http.Request)
}
