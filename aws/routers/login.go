package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Brackistar/golang-basic-backend/interfaces"
	"github.com/Brackistar/golang-basic-backend/shared/constants"
	"github.com/Brackistar/golang-basic-backend/shared/models"
	"github.com/Brackistar/golang-basic-backend/shared/utils"
	"github.com/aws/aws-lambda-go/events"
)

func Login(ctx *context.Context, responseBuilder interfaces.ResponseBuilder[events.APIGatewayProxyResponse]) *events.APIGatewayProxyResponse {
	log.Printf(routeBeginMsg, "Login")
	defer log.Print(routeEndMsg)

	responseBuilder.Clear()
	responseBuilder.SetStatusCode(http.StatusBadRequest)

	var user models.User
	body := utils.GetContextValue[string](ctx, constants.CtxKeyBdy)
	err := json.Unmarshal([]byte(body), &user)

	if err != nil {
		log.Printf(unmarshalErrorMsg, user, body)
		responseBuilder.SetBody(usrFailBodyMsg)
		return responseBuilder.Build()
	}

	if len(user.Email) == 0 {
		responseBuilder.SetBody(fmt.Sprintf(fldEmptyMsg, "Email"))
		return responseBuilder.Build()
	}

	if len(user.Pass) != constants.CtxPassLenght {
		responseBuilder.SetBody(usrPssInvalidMsg)
		return responseBuilder.Build()
	}

	client := getClient(ctx)

	savedUser, err := getUser(user.Email, client)

	if err != nil {
		responseBuilder.SetStatusCode(http.StatusNotFound)
		responseBuilder.SetBody(usrNotFoundMsg)
		return responseBuilder.Build()
	}

	if isUserValid(savedUser, user) {
		responseBuilder.Clear()
		responseBuilder.SetStatusCode(http.StatusOK)
		responseBuilder.SetBody(usrLoginSuccess)

		return responseBuilder.Build()
	}

	return responseBuilder.Build()
}

func isUserValid(savedUser models.User, loginUser models.User) bool {
	return savedUser.Email == loginUser.Email && savedUser.Pass == loginUser.Pass
}
