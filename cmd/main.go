package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/gleisonem/convert-html-pdf-golang/controllers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	cpu "github.com/shirou/gopsutil/cpu"
)

func main() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Printf("Memory usage before: %v MB\n", mem.Alloc/1024/1024)

	percent, _ := cpu.Percent(time.Second, false)
	fmt.Printf("CPU Load before: %v%%\n", percent[0])

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
