package main

import (
	"net/http"

	"github.com/schlucht/myGame/pkg/payment"
	"github.com/schlucht/myGame/server"
)

const PORT = "8080"

var pag server.Page

type AboutData struct {
	Name string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	res, _ := payment.ReadFile("./daten/pay.csv")
	pag = pag.New("Der Titel", "Die Erklärung", &res)
	server.RenderTemplate(w, "index.html", pag)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	dat := AboutData{
		Name: "Lothar",
	}
	pag = pag.New("About", "Email", &dat)
	server.RenderTemplate(w, "about.html", pag)
}

func main() {

	routes := []server.Route{
		{RouterName: "/", Handler: homeHandler},
		{RouterName: "/about", Handler: aboutHandler},
	}

	server := server.Server{
		Port:    PORT,
		Message: "Message",
		Static:  "/static/",
		Routes:  routes,
	}

	server.Init()
}
