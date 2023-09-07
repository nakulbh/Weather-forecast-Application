package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Loading environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Creating new Gorilla Mux router
	r := mux.NewRouter()

	// Defining endpoints
	r.HandleFunc("/current/{city}", getCurrentWeather).Methods("GET")

	// Enabling CORS middleware
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // You can specify specific origins if needed
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization"}),
	)

	// Wrapingrouter with the CORS middleware.
	r.Use(corsHandler)

	// Starting the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server is starting at port 8080 url := localhost:8080/current/{city}...")
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func getCurrentWeather(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	city := params["city"]
	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")

	// Constructing the URL for current weather data
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

	// Making a request to OpenWeatherMap API for current weather
	client := resty.New()
	resp, err := client.R().Get(url)

	if err != nil {
		http.Error(w, "Failed to retrieve weather data", http.StatusInternalServerError)
		return
	}

	if resp.StatusCode() != 200 {
		http.Error(w, "Failed to retrieve weather data", resp.StatusCode())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp.Body())
}

