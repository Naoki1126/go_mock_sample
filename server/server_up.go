package server

import (
	"log"
	"net/http"
)

func ServerUp() {
	log.Println(1111)
	http.HandleFunc("/api/get/test", func(w http.ResponseWriter, r *http.Request) {
		log.Print("aaaaa")
	})
	log.Fatal(http.ListenAndServe(":3000", nil))
}
