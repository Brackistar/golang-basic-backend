package handlers

import (
	"context"
	"log"

	"github.com/Brackistar/golang-basic-backend/interfaces"
	"github.com/Brackistar/golang-basic-backend/shared/constants"
	"github.com/Brackistar/golang-basic-backend/shared/utils"
	"github.com/aws/aws-lambda-go/events"
)

const (
	getHandlerBeginMsg string = "GET request to path \"%s\" being handled"
)

func handleGetRequest(ctx *context.Context, request *events.APIGatewayProxyRequest, responseBuilder interfaces.ResponseBuilder[events.APIGatewayProxyResponse]) *events.APIGatewayProxyResponse {
	log.Printf(getHandlerBeginMsg, utils.GetContextValue[string](ctx, constants.CtxKeyPath))

	return responseBuilder.Build()
}
