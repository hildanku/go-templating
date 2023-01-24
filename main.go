package main

import (
	"golang-slice/app/controllers/dashboardcontroller"
	"log"
	"net/http"
)

// penting
func dir() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("app/views/_layout/assets"))))
}

func main() {
	dir()
	// route dashboard
	http.HandleFunc("/", dashboardcontroller.Index)

	log.Println("Server running in 8080")
	http.ListenAndServe(":8000", nil)
}
