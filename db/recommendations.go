package db

import (
	"log"

	"stock-challenge/models"
)

func GetBestStocks(db *DBConnection) ([]models.Stock, error) {
	query := `
	SELECT ticker, company, brokerage, action, rating_from, rating_to, target_from, target_to, time
	FROM stocks
	WHERE (
		(rating_from = 'Sell' AND rating_to IN ('Neutral', 'Buy', 'Outperform')) OR
		(rating_from = 'Neutral' AND rating_to IN ('Buy', 'Outperform')) OR
		(rating_from = 'Underweight' AND rating_to IN ('Equal Weight', 'Overweight', 'Buy')) OR
		(rating_from = 'Equal Weight' AND rating_to IN ('Overweight', 'Buy')) OR
		(rating_from = 'Market Perform' AND rating_to IN ('Outperform', 'Buy')) OR
		(rating_from = 'Hold' AND rating_to IN ('Buy', 'Strong Buy')) OR
		(rating_from = 'Buy' AND rating_to IN ('Overweight', 'Buy', 'Strong Buy', 'Outperform'))
	)
	ORDER BY time DESC;
	`
	var stocks []models.Stock
	rows, err := db.Conn.Query(db.Ctx, query)
	if err != nil {
		log.Fatalf("Error fetching stocks: %v\n", err)
		return nil, err
	}
	defer rows.Close()

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
