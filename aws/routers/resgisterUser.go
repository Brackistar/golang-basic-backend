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

func RegisterUser(ctx *context.Context, responseBuilder interfaces.ResponseBuilder[events.APIGatewayProxyResponse]) *events.APIGatewayProxyResponse {
	log.Printf(routeBeginMsg, "RegisterUser")
	defer log.Println(routeEndMsg)

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

	if length := len(user.Pass); length != 64 {
		log.Printf("Tried user registration with password len: %d", length)
		responseBuilder.SetBody(fmt.Sprintf(invFieldValMsg, "Password", ""))
		return responseBuilder.Build()
	}

	client := getClient(ctx)

	if userExists(user.Email, client) {
		log.Printf("User with email: %s already exists", user.Email)

		responseBuilder.SetBody(usrRegExistsMsg)
		return responseBuilder.Build()
	}

	return createUser(user, client, responseBuilder)
}

func createUser(user models.User, client interfaces.DataOrigin, responseBuilder interfaces.ResponseBuilder[events.APIGatewayProxyResponse]) *events.APIGatewayProxyResponse {
	log.Printf("Start user creation for email: %s", user.Email)

	id, success, err := client.CreateRecord(constants.UsersOrigin, user)

	if err != nil {
		responseBuilder.SetBody(fmt.Sprintf(usrRegFailErrMsg, err))
		return responseBuilder.Build()
	}

	if !success {
		responseBuilder.SetBody(usrRegFailMsg)
		return responseBuilder.Build()
	}

	responseBuilder.SetStatusCode(http.StatusCreated)
	responseBuilder.SetBody(fmt.Sprintf(usrInserted, id.(string)))

	log.Print("User created")

	return responseBuilder.Build()
}

func userExists(email string, client interfaces.DataOrigin) bool {
	log.Printf("Checking for user pre-existance for email: %s", email)

	val, _ := client.GetRecord(constants.UsersOrigin, "email", email)

	return val != nil
}

func getClient(ctx *context.Context) interfaces.DataOrigin {
	dbOrigin := utils.GetContextValue[interfaces.DBConnectManager](ctx, constants.CtxKeyDbManager)
	return dbOrigin.GetDataOrigin()
}
