package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Go App is running!"))
	})

	http.HandleFunc("/db-test", func(w http.ResponseWriter, r *http.Request) {
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

		db, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Println("Error connecting to the database:", err)
			http.Error(w, "DB connection failed!", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		var now string
		err = db.QueryRow("SELECT NOW()").Scan(&now)
		if err != nil {
			log.Println("Error executing query:", err)
			http.Error(w, "DB query failed!", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("DB Test: " + now))
	})

	port := ":8080"
	log.Println("Starting server on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
