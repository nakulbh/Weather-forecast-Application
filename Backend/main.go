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

<<<<<<< HEAD
<<<<<<< HEAD
	// Enabling CORS middleware
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), 
=======
	// Enable CORS middleware
=======
	// Enabling CORS middleware
>>>>>>> 045bdae (making backend)
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // You can specify specific origins if needed
>>>>>>> 616c5ee (making backend)
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization"}),
	)

<<<<<<< HEAD
<<<<<<< HEAD
	// Wrapingrouter with the CORS middleware.
	r.Use(corsHandler)

	// Starting the server
=======
	// Wrap your router with the CORS middleware.
	r.Use(corsHandler)

	// Start the server
>>>>>>> 616c5ee (making backend)
=======
	// Wrapingrouter with the CORS middleware.
	r.Use(corsHandler)

	// Starting the server
>>>>>>> 045bdae (making backend)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
<<<<<<< HEAD
<<<<<<< HEAD
	fmt.Println("Server is starting at port 8080 url := localhost:8080/current/{city}...")
=======
	fmt.Println("Server is starting at port 8080...")
>>>>>>> 616c5ee (making backend)
=======
	fmt.Println("Server is starting at port 8080 url := localhost:8080/current/{city}...")
>>>>>>> 045bdae (making backend)
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

