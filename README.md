# Weather Forecast Web App

This is a simple web application that provides current weather data for a given city. The application consists of a backend API built with Go (Golang) and a frontend built with Html Javascript and BOOTSTRAP.

## Prerequisites

Before you start, make sure you have the following dependencies installed:

- Go (Golang)
- An OpenWeatherMap API key (sign up at https://openweathermap.org/ to get one)

## Backend Setup

1. **Clone this repository** to your local machine:

   ```bash
   git clone https://github.com/nakulbh/Weather-forecast-Application.git
   
2. **Navigate to the backend directory**
   '''bash
   cd Weather-forecast-Application/Backend

3. **In the .env file** in the backend directory add your OpenWeatherMap API key
  '''bash
  OPENWEATHERMAP_API_KEY=your-api-key-here
  PORT=8080

4. **Install Go dependencies**
  '''bash
  go mod tidy

5. **Start the backend server**
  '''bash
  go run main.go

  The backend API will be accessible at http://localhost:8080.

## Frontend Setup

1. **Navigate to the frontend directory**
   '''bash
   cd Weather-forecast-Application/Frontend

2. **Start the backend server**
  '''bash
  go run main.go

  The frontend will be accessible at http://localhost:5000.

## Usage

1. Open your web browser and navigate to http://localhost:5000 to access the Weather Forecast Web App.

2. Enter the name of a city in the input field and click the "Get Weather" button to retrieve weather information.

# Features
1. Get current weather data for a city.

# Technologies Used
1. Go (Golang) for the backend.
2. Html Bootstrap javascript for the frontend.
3. OpenWeatherMap API for weather data.




