package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gleisonem/convert-html-pdf-golang/services"
	"github.com/google/uuid"
)

func GenerateHtmlConverterHandler(w http.ResponseWriter, r *http.Request) {

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

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(response)
	} else {
		fmt.Println(err)
	}
}

func HtmlConverterHandler(w http.ResponseWriter, r *http.Request) {

	pdfService := services.NewRequestPdf("")

	templatePath := "templates/allreports.html"
	outputPath := "storage/allreports.pdf"

	templateData := struct {
		Title       string
		Description string
		Company     string
		Contact     string
		Country     string
	}{
		Title:       "HTML to PDF generator",
		Description: "This is the simple HTML to PDF file.",
		Company:     "Jhon Lewis",
		Contact:     "Maria Anders",
		Country:     "Germany",
	}

	if err := pdfService.ParseTemplate(templatePath, templateData); err == nil {
		args := []string{"grayscale"}

		ok, _ := pdfService.GeneratePDF(outputPath, args)
		fmt.Println(ok, "pdf generated successfully")

		response := map[string]string{
			"file":    outputPath,
			"message": "pdf generated successfully",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(response)
	} else {
		fmt.Println(err)
	}
}

func GenerateBufferHtmlConverterHandler(w http.ResponseWriter, r *http.Request) {

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
	outputPathFileName := uuid.New().String() + ".pdf"
	outputPath := folderStoragePath + outputPathFileName

	if err := pdfService.ParseToString(html.HtmlContentString); err == nil {
		args := []string{"grayscale"}

		_, err := pdfService.GeneratePDF(outputPath, args)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		f, err := os.Open(outputPath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		w.Header().Set("Content-Type", "application/pdf")

		io.Copy(w, f)
	} else {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}
