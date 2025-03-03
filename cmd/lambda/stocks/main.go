package main

import (
	stockslambda "stock-challenge/lambda/stocks"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(stockslambda.StocksLambdaHandler)
}
