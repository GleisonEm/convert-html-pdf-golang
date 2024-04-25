package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"

	"time"

	"github.com/gleisonem/convert-html-pdf-golang/services"
	"github.com/google/uuid"
)

func GenerateHtmlConverterHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Printf("Memory usage before route: %v MB\n", mem.Alloc/1024/1024)

	cpu := runtime.NumCPU()
	fmt.Printf("CPU Load before route: %v%%\n", cpu)

	pdfService := services.NewRequestPdf("")

	var html struct {
		HtmlContentString string `json:"html_content"`
	}

	err := json.NewDecoder(r.Body).Decode(&html)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if html.HtmlContentString == "" {
		http.Error(w, "O campo html_content não pode ser nulo", http.StatusBadRequest)
		return
	}

	folderStoragePath := "storage/"
	folderStaticPath := "public/"
	outputPathFileName := uuid.New().String() + ".pdf"
	outputPath := folderStoragePath + outputPathFileName

	if err := pdfService.ParseToString(html.HtmlContentString); err == nil {
		args := []string{"grayscale"}

		_, err := pdfService.GeneratePDF(outputPath, args)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := map[string]string{
			"file":     folderStaticPath + outputPathFileName,
			"filename": outputPathFileName,
			"message":  "pdf generated successfully",
		}

		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		fmt.Printf("Memory usage after: %v MB\n", mem.Alloc/1024/1024)

		cpu := runtime.NumCPU()
		fmt.Printf("CPU Load after: %v%%\n", cpu)

		elapsed := time.Since(start)
		fmt.Printf("Route executed in %s\n", elapsed)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(response)
	} else {
		fmt.Println(err)
	}
}
