package api

import (
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
	conn, ctx := db.ConnectDB(databaseURL)
	defer conn.Close(ctx)

	stocks, err := fetchStocks()
	if err != nil {
		fmt.Println("Error fetching stocks: ", err)
		return
	}

	for _, stock := range stocks {
		db.InsertStocks(conn, ctx, stock)
	}

	fmt.Println("Stocks inserted successfully")
}
