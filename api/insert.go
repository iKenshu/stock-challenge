package api

import (
	"fmt"

	"stock-challenge/db"
	"stock-challenge/models"
)

func InsertStocksFromAPI(fetchStocks func() ([]models.Stock, error)) {
	conn, ctx := db.ConnectDB()
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
