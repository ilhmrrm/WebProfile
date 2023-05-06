package main

import (
	"WebProfile/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.DashboardHandler)

	fileserver := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	log.Println("Server running at port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
