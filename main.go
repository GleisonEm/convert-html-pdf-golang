package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gleisonem/convert-html-pdf-golang/controllers"
	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8087"
	}

	r := mux.NewRouter()
	r.HandleFunc("/generate", controllers.HtmlConverterHandler).Methods("GET")
	http.Handle("/", r)

	fs := http.FileServer(http.Dir("./storage/"))
    r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	address := ":" + port
	fmt.Printf("Servidor escutando na porta %s...\n", port)
	http.ListenAndServe(address, nil)
}
