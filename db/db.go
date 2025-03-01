package db

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"stock-challenge/models"

	"github.com/jackc/pgx/v4"
)

func ConnectDB(databaseURL string) (*pgx.Conn, context.Context) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	return conn, ctx
}

func InsertStocks(conn *pgx.Conn, ctx context.Context, stock models.Stock) {
	targetFrom, err := parsePrice(stock.TargetFrom)
	if err != nil {
		fmt.Println("Error parsing price: ", err)
	}
	targetTo, err := parsePrice(stock.TargetTo)
	if err != nil {
		fmt.Println("Error parsing price: ", err)
	}

	query := `INSERT INTO stocks (ticker, company, brokerage, action, rating_from, rating_to, target_from, target_to, time)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
	ON CONFLICT (ticker) DO UPDATE SET
	action = EXCLUDED.action,
	rating_from = EXCLUDED.rating_from,
	rating_to = EXCLUDED.rating_to,
	target_from = EXCLUDED.target_from,
	target_to = EXCLUDED.target_to;`
	_, err = conn.Exec(
		ctx,
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
		fmt.Println("Error inserting stock: ", err)
	}
}

func FetchAllStocks(conn *pgx.Conn, ctx context.Context) ([]models.Stock, error) {
	query := `SELECT ticker, company, brokerage, action, rating_from, rating_to, target_from, target_to, time FROM stocks;`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		fmt.Println("Error fetching stocks: ", err)
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
			fmt.Println("Error scanning stock: ", err)
			return nil, err
		}
		stocks = append(stocks, stock)
	}

	return stocks, nil
}

func GetBestStocks(conn *pgx.Conn, ctx context.Context) ([]models.Stock, error) {
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
	rows, err := conn.Query(ctx, query)
	if err != nil {
		fmt.Println("Error fetching stocks: ", err)
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
			fmt.Println("Error scanning stock: ", err)
			return nil, err
		}
		stocks = append(stocks, stock)
	}

	return stocks, nil
}

func parsePrice(priceStr string) (float64, error) {
	priceStr = strings.TrimPrefix(priceStr, "$")
	return strconv.ParseFloat(priceStr, 64)
}
