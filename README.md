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

5 **Start the backend server**
  '''bash
  go run main.go

  The backend API will be accessible at http://localhost:8080.




