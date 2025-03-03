package api

import (
	"context"
	"fmt"
	"os"

	"stock-challenge/db"
	"stock-challenge/models"

	"github.com/joho/godotenv"
)

func InsertStocksFromAPI(fetchStocks func() ([]models.Stock, error)) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	dbConn := db.ConnectDB(databaseURL)
	defer dbConn.Close(context.Background())

	stocks, err := fetchStocks()
	if err != nil {
		fmt.Println("Error fetching stocks: ", err)
		return
	}

	for _, stock := range stocks {
		db.InsertStocks(dbConn, stock)
	}

	fmt.Println("Stocks inserted successfully")
}
