package main

import (
	"ascii-art/functions"
	"html/template"
	"net/http"
	"fmt"
	"strings"
)

type Data struct {
	Str string
	Banner string
	Res string
	A	template.HTML
}
var  data Data
func processHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if  r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	
	data.Str = r.FormValue("data")
	data.Banner = r.FormValue("banner")
	if !function.CheckBanner(data.Banner) {
		http.Error(w, "Invalid banner", http.StatusBadRequest)
		return
	}

	data.Str = strings.ReplaceAll(data.Str, "\r\n", "\\n")
	data.Str = strings.ReplaceAll(data.Str, "\n", "\\n")
	data.Res = function.TraitmentData(data.Banner, data.Str)
	data.A = template.HTML(strings.ReplaceAll(data.Res, "\n", "<br>"))
	if err := t.Execute(w, data); err != nil {
		http.Error(w,  err.Error(), http.StatusInternalServerError)
		return
	}
}



func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusInternalServerError)
		return
	}
	t, err := template.ParseFiles("home.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := t.Execute(w, nil); err != nil {
		http.Error(w, err.Error(),  http.StatusInternalServerError)
		return
	}
}

func exportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if data.Res != "" {
		// fmt.Println(data.Res)
		w.Header().Set("Content-Disposition", "attachment; filename=ascii-art.txt")
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Length",  fmt.Sprintf("%d", len(data.Res)))
		if _, err := w.Write([]byte(data.Res)); err != nil {
			http.Error(w,  err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "No data to export", http.StatusNotFound)
		return
	}
}


func main() {
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css", fs))
	http.HandleFunc("/", homePage)
	http.HandleFunc("/process", processHandler)
	http.HandleFunc("/export", exportHandler)
	fmt.Println("Server is running at http://localhost:8082")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}