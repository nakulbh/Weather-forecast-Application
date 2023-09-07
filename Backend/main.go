package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := mux.NewRouter()

	router.HandleFunc("/current/{city}", getCurrentWeather).Methods("Get")
	router.HandleFunc("/forecast/{city}", getWeatherForecast).Methods("Get")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

func getCurrentWeather(w http.ResponseWriter, r *http.Request)

func getWeatherForecast(w http.ResponseWriter, r *http.Request) {

}
