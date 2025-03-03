package recommendationslambda

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"stock-challenge/db"

	"github.com/aws/aws-lambda-go/events"
)

func RecommendationsLambdaHandler(ctx context.Context) (events.APIGatewayProxyResponse, error) {
	databaseURL := os.Getenv("DATABASE_URL")

	dbConn := db.ConnectDB(databaseURL)
	defer dbConn.Close(ctx)

	recommendations, err := db.GetBestStocks(dbConn)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"Content-Type":                 "application/json",
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Methods": "GET",
			},
			Body: err.Error(),
		}, nil
	}

	body, err := json.Marshal(recommendations)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"Content-Type":                 "application/json",
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Methods": "GET",
			},
			Body: err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "GET",
		},
		Body: string(body),
	}, nil
}
