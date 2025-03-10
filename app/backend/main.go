package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"fmt"
	"net/url"
  "github.com/Azure/go-autorest/autorest/adal"

	_ "github.com/go-sql-driver/mysql"
)

type Patient struct {
	ID         int    `json:"id"`
	FullName   string `json:"full_name"`
	Department string `json:"department"`
	BedNumber  int    `json:"bed_number"`
}

var db *sql.DB

func getAccessToken() (string, error) {
	msiEndpoint := "http://169.254.169.254/metadata/identity/oauth2/token"
	resource := "https://ossrdbms-aad.database.windows.net/"

	// Token holen von Azure Managed Identity
	token, err := adal.NewServicePrincipalTokenFromMSI(msiEndpoint, resource)
	if err != nil {
		return "", fmt.Errorf("Fehler beim Abrufen des Tokens: %v", err)
	}

	err = token.Refresh()
	if err != nil {
		return "", fmt.Errorf("Fehler beim Aktualisieren des Tokens: %v", err)
	}

	return token.OAuthToken(), nil
}

func initDB() {
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	var err error
	token, err := getAccessToken()
	if err != nil {
    log.Fatal(err)
  }
	dsn := fmt.Sprintf("%s@tcp(%s:3306)/%s?tls=true",
		url.QueryEscape(token), dbHost, dbName)

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS patient (id INT AUTO_INCREMENT PRIMARY KEY, full_name TEXT, department TEXT, bed_number INT)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
}

func getPatients(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, full_name, department, bed_number FROM patient")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	patients := []Patient{}
	for rows.Next() {
		var patient Patient
		rows.Scan(&patient.ID, &patient.FullName, &patient.Department, &patient.BedNumber)
		patients = append(patients, patient)
	}
	json.NewEncoder(w).Encode(patients)
}

func addPatient(w http.ResponseWriter, r *http.Request) {
	var patient Patient
	json.NewDecoder(r.Body).Decode(&patient)
	statement, err := db.Prepare("INSERT INTO patient (full_name, department, bed_number) VALUES (?, ?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	statement.Exec(patient.FullName, patient.Department, patient.BedNumber)
	w.WriteHeader(http.StatusCreated)
}

func deletePatient(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	statement, err := db.Prepare("DELETE FROM patient WHERE id = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = statement.Exec(idParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	initDB()
	http.HandleFunc("/patients", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getPatients(w, r)
		} else if r.Method == http.MethodPost {
			addPatient(w, r)
		} else if r.Method == http.MethodDelete {
			deletePatient(w, r)
		}
	})

	log.Println("Backend server started on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
