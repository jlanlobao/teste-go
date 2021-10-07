package main

//package main

import (
	"net/http"

	"github.com/jlobao-trn/alura2/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
