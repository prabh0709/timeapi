package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Initialize the database connection
func init() {
	var err error
	// Update with your MySQL credentials
	dsn := "root:root@tcp(127.0.0.1:3306)/toronto_time"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
}

// Get the current time in Toronto
func getTorontoTime() (time.Time, error) {
	location, err := time.LoadLocation("America/Toronto")
	if err != nil {
		return time.Time{}, err
	}
	return time.Now().In(location), nil
}

// Log the time to the MySQL database
func logTimeToDatabase(t time.Time) {
	query := "INSERT INTO time_log (timestamp) VALUES (?)"
	_, err := db.Exec(query, t)
	if err != nil {
		fmt.Println("Error logging time:", err)
	}
}

// Handle the /current-time API request
func handleCurrentTime(w http.ResponseWriter, r *http.Request) {
	torontoTime, err := getTorontoTime()
	if err != nil {
		http.Error(w, "Unable to get Toronto time", http.StatusInternalServerError)
		return
	}

	// Return the current time in JSON format
	response := map[string]string{
		"current_time": torontoTime.Format("2006-01-02 15:04:05"),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	// Log the time to MySQL
	logTimeToDatabase(torontoTime)
}

func main() {
	http.HandleFunc("/current-time", handleCurrentTime)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
