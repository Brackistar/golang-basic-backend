package handlers

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/Brackistar/golang-basic-backend/interfaces"
	"github.com/Brackistar/golang-basic-backend/jwt"
	"github.com/Brackistar/golang-basic-backend/shared/constants"
	"github.com/Brackistar/golang-basic-backend/shared/models"
	"github.com/Brackistar/golang-basic-backend/shared/utils"
	"github.com/aws/aws-lambda-go/events"
)

const (
	nonAuthPaths  string = "register|login|avatar|banner"
	authHeaderKey string = "Authorization"
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

	isOk, statusCode, msg, _ := authorize(ctx, request)

	if !isOk {
		responseBuilder.SetStatusCode(uint(statusCode))
		responseBuilder.SetBody(msg)

		return responseBuilder.Build()
	}

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

func authorize(ctx *context.Context, request *events.APIGatewayProxyRequest) (bool, int, string, *models.Claim) {
	log.Println("Autorizing token")

	path := utils.GetContextValue[string](ctx, constants.CtxKeyPath)

	if nonAuth(path) {
		return true, http.StatusOK, "", &models.Claim{}
	}

	token := request.Headers[authHeaderKey]

	invalidAuthMsg := func(msg string, status int) (bool, int, string, *models.Claim) {
		return false, status, msg, &models.Claim{}
	}

	if len(token) == 0 {
		return invalidAuthMsg(noTokenMsg, http.StatusUnauthorized)
	}

	claim, ok, msg, err := jwt.HandleToken(token, utils.GetContextValue[string](ctx, constants.CtxKeyJwt))

	if err != nil {
		log.Printf(unexErrAuthMsg, err)
		return invalidAuthMsg(failAuthMsg, http.StatusUnauthorized)
	}

	if !ok {
		log.Printf(msg, err)
		return invalidAuthMsg(msg, http.StatusUnauthorized)
	}

	log.Println("Token auth ok")

	return true, http.StatusOK, msg, claim
}

func nonAuth(path string) bool {
	paths := strings.Split(nonAuthPaths, "|")
	var result bool = true

	for _, p := range paths {
		result = path == p
	}

	return result
}
