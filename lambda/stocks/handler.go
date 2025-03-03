package stockslambda

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"stock-challenge/db"

	"github.com/aws/aws-lambda-go/events"
)

func StocksLambdaHandler(ctx context.Context) (events.APIGatewayProxyResponse, error) {
	databaseURL := os.Getenv("DATABASE_URL")

	dbConn := db.ConnectDB(databaseURL)
	defer dbConn.Close(ctx)

	stocks, err := db.FetchAllStocks(dbConn)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: err.Error(),
		}, nil
	}

	body, err := json.Marshal(stocks)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, OPTIONS",
			"Access-Control-Allow-Headers": "Content-Type, Authorization, X-Requested-With",
		},
		Body: string(body),
	}, nil
}
