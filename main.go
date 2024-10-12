package main

import (
	"ascii-art/functions"
	"fmt"
	"html/template"
	"net/http"

	"strings"
)


type Data struct {
	Str string
	Banner string
	Res string
	A	template.HTML
}


func processHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w,  "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}
	var data Data
	data.Str = r.FormValue("data")
	if len(data.Str) > 200 {
		http.Error(w, "Input data exceeds 200 characters limit.", http.StatusBadRequest)
		return
	}

	data.Banner = r.FormValue("banner")
	if !function.BannerExists(data.Banner) {
		http.Error(w, "Banner not found", http.StatusNotFound)
		return
	}

	data.Str = strings.ReplaceAll(data.Str, "\r\n", "\n")
	
	data.Res = function.TraitmentData(data.Banner, data.Str)
	if data.Res == "" { 
		http.Error(w, "Internal Server Error: Failed to process data.", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename=ascii-art.txt")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length",  fmt.Sprintf("%d", len(data.Res)))
	if _, err := w.Write([]byte(data.Res)); err != nil {
		http.Error(w,  err.Error(), http.StatusInternalServerError)
		return
	}
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Error 404 : Not Found", http.StatusNotFound)
		return
	}
	t, err := template.ParseFiles("home.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	if err := t.Execute(w, nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
func main() {

	http.HandleFunc("/", pageHandler)
	http.HandleFunc("/ascii-art", processHandler)
	fmt.Println("Server is running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}