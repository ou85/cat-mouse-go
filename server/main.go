package main

import (
	"log"
	"net/http"
)

func main() {
	// fs := http.FileServer(http.Dir("../static"))
	fs := http.FileServer(http.Dir("../"))
	http.Handle("/", fs)

	log.Println("========================================")
	log.Println("🚀 Server is up and running!")
	log.Println("🌐 Access it at: http://localhost:8080")
	log.Println("========================================")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
