package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/weather", getWeather).Methods("Get")
	err := http.ListenAndServe(":8080", router)
	fmt.Printf("Starting the server at port 8080 ..")
	if err != nil {
		log.Fatal(err)
	}

}

func getWeather(w http.ResponseWriter, r *http.Request) {

}
