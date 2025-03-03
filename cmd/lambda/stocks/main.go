package main

import (
	stocksLambda "stock-challenge/lambda/stocks"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(stocksLambda.StocksLambdaHandler)
}
