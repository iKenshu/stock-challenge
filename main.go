// This file do a requests endpoint API to get the data from the API

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"net/http"
	"stock-challenge/api"
	"stock-challenge/db"
	"stock-challenge/models"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	apiURL := os.Getenv("API_URL")
	apiToken := os.Getenv("API_TOKEN")
	databaseURL := os.Getenv("DATABASE_URL")

	conn, ctx := db.ConnectDB(databaseURL)
	defer conn.Close(ctx)

	fmt.Println("Inserting stocks from API")
	api.InsertStocksFromAPI(func() (stocks []models.Stock, err error) {
		return api.FetchStocks(apiURL, apiToken)
	})

	http.Handle(
		"/api/stocks",
		corsMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			conn, ctx := db.ConnectDB()
			defer conn.Close(ctx)

			stocks, err := db.FetchAllStocks(conn, ctx)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(stocks)
		})),
	)

	http.Handle(
		"/api/recommendations",
		corsMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			conn, ctx := db.ConnectDB()
			defer conn.Close(ctx)

			recommendations, err := db.GetBestStocks(conn, ctx)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(recommendations)
		})),
	)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
