# AwsGolangLambda
An AWS lambda written in golang to exercise with an API Gateway.

GOOS=linux go build -o main ./awsGolangLambda.go

build-lambda-zip -o main.zip main

Upload as function code to AWS lambda service and run with an API Gateway

As instructed by https://www.youtube.com/watch?v=SmNLIY9j3Ls
