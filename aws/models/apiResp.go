package models

import (
	"github.com/aws/aws-lambda-go/events"
)

type ApiResp struct {
	Status         int
	Message        string
	CustomResponse *events.APIGatewayProxyRequest
}
