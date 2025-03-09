package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	if host == "" || user == "" || password == "" || database == "" {
		log.Fatal("Missing required database environment variables")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database)
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Database not reachable: %v", err)
	}

	log.Println("Connected to database successfully")
}

func handler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT NOW()")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var now string
	if rows.Next() {
		if err := rows.Scan(&now); err != nil {
			http.Error(w, "Error scanning result", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Database Time: " + now))
}

func main() {
	initDB()
	defer db.Close()

	http.HandleFunc("/", handler)
	log.Println("Server started on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
