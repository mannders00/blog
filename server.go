package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", homeHandler).Methods("GET")

	router.HandleFunc("/blog", getBlogHandler).Methods("GET")
	router.HandleFunc("/blog/{blogID}", blogHandler).Methods("GET")

	router.HandleFunc("/project/{projectID}", projectHandler).Methods("GET")

	router.HandleFunc("/book", bookHandler).Methods("GET")

	router.PathPrefix("/public/").HandlerFunc(getPublic)

	fmt.Println("starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("public/templates/header.tmpl", "public/templates/footer.tmpl", "public/html/index.html"))
	err := tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getBlogHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("public/templates/header.tmpl", "public/templates/footer.tmpl", "public/html/blog.html"))
	err := tmpl.ExecuteTemplate(w, "blog.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("public/templates/header.tmpl", "public/templates/footer.tmpl", "public/html/book.html"))
	err := tmpl.ExecuteTemplate(w, "book.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blogID := vars["blogID"]

	tmpl := template.Must(template.ParseFiles("public/templates/header.tmpl", "public/templates/footer.tmpl", fmt.Sprintf("public/html/blog/%s.html", blogID)))
	err := tmpl.ExecuteTemplate(w, fmt.Sprintf("%s.html", blogID), nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func projectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectID := vars["projectID"]

	tmpl := template.Must(template.ParseFiles("public/templates/header.tmpl", "public/templates/footer.tmpl", fmt.Sprintf("public/html/projects/%s.html", projectID)))
	err := tmpl.ExecuteTemplate(w, fmt.Sprintf("%s.html", projectID), nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getPublic(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	filePath := filepath.Join(".", path)

	if strings.HasPrefix(path, "/public/css") {
		w.Header().Set("Content-Type", "text/css")
	} else if strings.HasPrefix(path, "/public/js") {
		w.Header().Set("Content-Type", "application/javascript")
	} else if strings.HasPrefix(path, "/public/images") {
		w.Header().Set("Content-Type", "image/jpeg")
	} else if strings.HasPrefix(path, "/public/images") {
		w.Header().Set("Content-Type", "image/png")
	} else if strings.HasPrefix(path, "/public/fonts") {
		w.Header().Set("Content-Type", "fonts/font")
	}

	http.ServeFile(w, r, filePath)
}
