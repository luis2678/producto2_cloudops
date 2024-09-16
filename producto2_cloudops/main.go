package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body><img src='/static/logo uoc.png' alt='UOC'><h1 style='color: blue'>Soy alumno de la UOC</h1></body></html>")

}

func main() {
	// Servir archivos estáticos
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handler)

	fmt.Println("App Server in a port:8080")
	http.ListenAndServe(":8080", nil)
}
