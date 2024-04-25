package main

import (
	"net/http"
	"github.com/gabriel/loja/routes"
)

func main() {
	routes.Rotas()
	http.ListenAndServe(":3000", nil)
}

