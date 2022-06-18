package main

import (
	"go_mock_sample/server"
	"log"
	"net/http"
)

func main() {
	log.Println(1111)
	http.HandleFunc("/api/get/task", func(w http.ResponseWriter, r *http.Request) {
		b := server.RequestGet("http://localhost:8080/api/get/task")
		w.Write(b)
	})

	http.HandleFunc("/api/get/book", func(w http.ResponseWriter, r *http.Request) {
		b := server.RequestGet("http://localhost:8080/api/get/book")
		w.Write(b)
	})
	log.Fatal(http.ListenAndServe(":3000", nil))
}
