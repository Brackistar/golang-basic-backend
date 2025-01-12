package main

import (
	"github.com/Brackistar/golang-basic-backend/aws"
	"github.com/Brackistar/golang-basic-backend/aws/configmanager"
	"github.com/Brackistar/golang-basic-backend/aws/responsebuilder"
	"github.com/Brackistar/golang-basic-backend/aws/secretsmanager"
	"github.com/Brackistar/golang-basic-backend/db"

	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	handler := &aws.AWSRequestHandler{
		ResponseBuilder: responsebuilder.NewAWSResponseBuilder(),
		SecretsManager:  secretsmanager.NewAWSSecretsManager(),
		ConfigManager:   configmanager.NewAwsConfigManager(),
		DbManager:       db.NewMongoConnectManager(),
	}

	lambda.Start(handler.HandleRequest)
}
