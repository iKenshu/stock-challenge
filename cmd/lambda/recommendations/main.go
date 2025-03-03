package main

import (
	recommendationslambda "stock-challenge/lambda/recommendations"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(recommendationslambda.RecommendationsLambdaHandler)
}
