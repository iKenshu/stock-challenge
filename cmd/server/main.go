package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"stock-challenge/api"
	"stock-challenge/db"
	"stock-challenge/handlers"
	"stock-challenge/middleware"
	"stock-challenge/models"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	apiURL := os.Getenv("API_URL")
	apiToken := os.Getenv("API_TOKEN")
	databaseURL := os.Getenv("DATABASE_URL")

	dbConn := db.ConnectDB(databaseURL)
	defer dbConn.Close(context.Background())

	log.Println("Starting to insert stocks from API")
	api.InsertStocksFromAPI(func() (stocks []models.Stock, err error) {
		return api.FetchStocks(apiURL, apiToken)
	})

	log.Println("Starting server on port 8080")
	http.Handle(
		"/stocks",
		middleware.CorsMiddleware(handlers.StocksHandler(dbConn)),
	)

	http.Handle(
		"/recommendations",
		middleware.CorsMiddleware(handlers.RecommendationsHandler(dbConn)),
	)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
