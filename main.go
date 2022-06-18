package main

import (
	"encoding/json"
	"fmt"
	"go_mock_sample/object"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/api/get/task", GetTaskHandler)
	http.HandleFunc("/api/get/book", GetBookHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func CommonHandler(w http.ResponseWriter, r *http.Request, i interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	j, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling", err)
		return
	}

	w.Write(j)

}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GetTaskHandler start")
	t := object.Task{}
	t.Init()
	fmt.Println(t)
	CommonHandler(w, r, t)
}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GetBookHandler start")
	b := object.Book{}
	b.Init()
	fmt.Println(b)
	CommonHandler(w, r, b)
}
