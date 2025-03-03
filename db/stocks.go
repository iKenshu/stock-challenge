package db

import (
	"log"

	"stock-challenge/models"
)

func InsertStocks(db *DBConnection, stock models.Stock) {
	targetFrom, err := parsePrice(stock.TargetFrom)
	if err != nil {
		log.Fatalf("Error parsing price: %v\n", err)
	}
	targetTo, err := parsePrice(stock.TargetTo)
	if err != nil {
		log.Fatalf("Error parsing price: %v\n", err)
	}

	query := `
	INSERT INTO stocks (ticker, company, brokerage, action, rating_from, rating_to, target_from, target_to, time)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
	ON CONFLICT (ticker) DO UPDATE SET
		action = EXCLUDED.action,
		rating_from = EXCLUDED.rating_from,
		rating_to = EXCLUDED.rating_to,
		target_from = EXCLUDED.target_from,
		target_to = EXCLUDED.target_to;
	`
	_, err = db.Conn.Exec(
		db.Ctx,
		query,
		stock.Ticker,
		stock.Company,
		stock.Brokerage,
		stock.Action,
		stock.RatingFrom,
		stock.RatingTo,
		targetFrom,
		targetTo,
		stock.Time,
	)
	if err != nil {
		log.Fatalf("Error inserting stock: %v\n", err)
	}
}

func FetchAllStocks(db *DBConnection) ([]models.Stock, error) {
	query := `
	SELECT ticker, company, brokerage, action, rating_from, rating_to, target_from, target_to, time
	FROM stocks;`

	rows, err := db.Conn.Query(db.Ctx, query)
	if err != nil {
		log.Fatalf("Error fetching stocks: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var stocks []models.Stock
	for rows.Next() {
		var stock models.Stock
		err := rows.Scan(
			&stock.Ticker,
			&stock.Company,
			&stock.Brokerage,
			&stock.Action,
			&stock.RatingFrom,
			&stock.RatingTo,
			&stock.TargetFrom,
			&stock.TargetTo,
			&stock.Time,
		)
		if err != nil {
			log.Fatalf("Error scanning stock: %v\n", err)
			return nil, err
		}
		stocks = append(stocks, stock)
	}

	return stocks, nil
}
