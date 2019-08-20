# AwsGolangLambda
An AWS lambda written in golang to exercise with an API Gateway.

GOOS=linux go build -o main ./awsGolangLambda.go
build-lambda-zip -o main.zip main
Upload to AWS and run with an API Gateway
