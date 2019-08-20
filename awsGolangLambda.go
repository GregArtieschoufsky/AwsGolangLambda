// Lambda Function for API Gateway
// 
// Notes for building and running
// GOOS=linux go build -o main ./awsGolangLambda.go
// build-lambda-zip -o main.zip main
// Upload to AWS and run with an API Gateway
package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"errors"
)

func main() {
	lambda.Start(HandleRequest)
}

// HandleRequest handles the request
func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod == "GET" {
		var stringResponse = "Yay a successful Response!!"
		APIResponse := events.APIGatewayProxyResponse{Body: stringResponse, StatusCode: 200}
		return APIResponse, nil
	}

	err := errors.New("method not allowed")
	APIResponse := events.APIGatewayProxyResponse{Body: "Method Not OK", StatusCode: 502}
	return APIResponse, err
}