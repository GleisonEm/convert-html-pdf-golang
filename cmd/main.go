package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gleisonem/convert-html-pdf-golang/controllers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8089"
	}

	r := mux.NewRouter()
	r.HandleFunc("/generate", controllers.GenerateHtmlConverterHandler).Methods("POST")
	http.Handle("/", r)

	fs := http.FileServer(http.Dir("./storage/"))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	address := ":" + port
	fmt.Printf("Servidor escutando na porta %s...\n", port)
	http.ListenAndServe(address, nil)
}
