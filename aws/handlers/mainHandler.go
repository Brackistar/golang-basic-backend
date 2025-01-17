package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/Brackistar/golang-basic-backend/interfaces"
	"github.com/Brackistar/golang-basic-backend/shared/constants"
	"github.com/Brackistar/golang-basic-backend/shared/utils"
	"github.com/aws/aws-lambda-go/events"
)

const (
	handleBeginMsg string = "Handling request %s > %s"
	badRqstMsg     string = "An error was found, and your request could not be processed"
	rqstEndMsg     string = "Request handled"
	ftlErrMsg      string = "Fatal error on method %s, path: %s, error message: %s"
	intErrMsg      string = "An unexpected error occurred, please try again later"
)

/*
Handles an AWS request of any method and path, redirecting the request to the correct inner handler based on the HTTP Method used
Allows handling requests for methods POST, GET, PUT and DELETE
In case of a internal panic, this function will gracefully handle it to generate an HTTP Internal Server Error response

ctx *context.Context: Pointer to the main context of the function
request *events.APIGatewayProxyRequest: Pointer to AWS API Gateway request to handle
responseBuilder  interfaces.ResponseBuilder[events.APIGatewayProxyResponse]: Custom builder that allows for dynamic response building

returns: response *events.APIGatewayProxyResponse
*/
func HandleRequest(ctx *context.Context, request *events.APIGatewayProxyRequest, responseBuilder interfaces.ResponseBuilder[events.APIGatewayProxyResponse]) (response *events.APIGatewayProxyResponse) {
	log.Printf(handleBeginMsg, utils.GetContextValue[string](ctx, constants.CtxKeyPath), utils.GetContextValue[string](ctx, constants.CtxKeyMethod))

	defer func() {

		if r := recover(); r != nil {
			log.Printf(ftlErrMsg, request.HTTPMethod, request.Path, r)

			responseBuilder.Clear()
			responseBuilder.SetStatusCode(http.StatusInternalServerError)
			responseBuilder.SetBody(intErrMsg)

			response = responseBuilder.Build()
		}

		log.Print(rqstEndMsg)
	}()

	responseBuilder.SetStatusCode(http.StatusBadRequest)

	switch utils.GetContextValue[string](ctx, constants.CtxKeyMethod) {
	case "POST":
		return handlePostRequest(ctx, request, responseBuilder)
	case "GET":
		return handleGetRequest(ctx, request, responseBuilder)
	case "PUT":
		return handlePutRequest(ctx, request, responseBuilder)
	case "DELETE":
		return handleDeleteRequest(ctx, request, responseBuilder)
	default:
		responseBuilder.SetBody(badRqstMsg)
		response = responseBuilder.Build()
	}

	return response
}
