package handlers

import (
	"encoding/json"
	"net/http"

	"stock-challenge/db"
)

func RecommendationsHandler(conn *db.DBConnection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		recommendations, err := db.GetBestStocks(conn)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(recommendations)
	}
}
