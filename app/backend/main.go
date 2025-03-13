package main

import (
  "context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"fmt"
  "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
  "github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"

	_ "github.com/go-sql-driver/mysql"
)

type Patient struct {
	ID         int    `json:"id"`
	FullName   string `json:"full_name"`
	Department string `json:"department"`
	BedNumber  int    `json:"bed_number"`
}

var db *sql.DB

func getAccessToken(client_id string) (string, error) {
  clientID := azidentity.ClientID(client_id)
	opts := azidentity.ManagedIdentityCredentialOptions{ID: clientID}
  cred, err := azidentity.NewManagedIdentityCredential(&opts)
  if err != nil {
      log.Fatalf("Failed to create default credential: %v", err)
  }

  // Create a token acquisition context
  ctx := context.Background()
  token, err := cred.GetToken(ctx, policy.TokenRequestOptions{
      Scopes: []string{"https://ossrdbms-aad.database.windows.net"},
  })
  if err != nil {
      log.Fatalf("Failed to get token: %v", err)
  }

	return token.Token, nil
}

func initDB() {
  dbUsername := os.Getenv("DB_USERNAME")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
  client_id := os.Getenv("AZURE_CLIENT_ID")

	var err error
	token, err := getAccessToken(client_id)
	if err != nil {
    log.Fatal(err)
  }
  fmt.Println(token)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=true&allowCleartextPasswords=true",
	    dbUsername,
	    token,
  		dbHost,
  		dbName,
  )

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
