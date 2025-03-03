package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"stock-challenge/db"
)

func StocksHandler(conn *db.DBConnection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stock, err := db.FetchAllStocks(conn)
		if err != nil {
			log.Fatalf("Error fetching stocks: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(stock)
	}
}
