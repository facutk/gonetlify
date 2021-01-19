package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/joho/godotenv/autoload"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Println("hi")
	params := request.QueryStringParameters

	name := "World"
	if params["name"] != "" {
		name = params["name"]
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hi " + name,
	}, nil

}

func main() {
	lambda.Start(handler)
}
