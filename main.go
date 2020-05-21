package main

import (
	"log"
	"net/http"
	"sbdb-semester/handler"
)

func main() {
	//http.HandleFunc goes here
	http.HandleFunc("/ping", handler.PingPongHandler)
	http.HandleFunc("/semester", handler.Handler)
	http.HandleFunc("/semesters", handler.AllHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
