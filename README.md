# Go Time Management Application
This project is a simple Go-based web application that connects to a MySQL database and provides functionality for managing and retrieving time-related data, specifically for Toronto.

# Features
- **Database Connection**: Establishes a connection to a MySQL database.
- **Toronto Time Retrieval**: Retrieves the current time for the Toronto timezone.
- **Web Server**: Handles HTTP requests for time-related operations.
  
# Requirements
- **Go Programming Language**: Version 1.17 or later.
- **MySQL Database**: A running instance of MySQL.
- **Go Modules**: Ensure all dependencies are installed via Go modules.

# Setup
1. Configure MySQL Database
1.	Create a MySQL database named toronto_time.
2.	Update the DSN (Data Source Name) in the main.go file with your MySQL credentials. Example:
dsn := "your_username:your_password@tcp(127.0.0.1:3306)/toronto_time"

# Install Dependencies
Run the following command to install dependencies:
go mod tidy

# Run the Application
Use the command below to start the server:
go run main.go
Endpoints
1. Get Current Time in Toronto
- **URL**: /getTorontoTime
- **Method**: GET
- **Description**: Fetches and returns the current time in the Toronto timezone.
   
# Example Usage
1.	Start the server using go run main.go.
2.	Open your browser or use a tool like Postman to visit http://localhost:8080/current-Time.

