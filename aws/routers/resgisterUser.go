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

	if len(user.Pass) != 128 {
		responseBuilder.SetBody(fmt.Sprintf(invFieldValMsg, "Password", ""))
		return responseBuilder.Build()
	}

	client := getClient(ctx)

	if userExists(user.Email, client) {
		responseBuilder.SetBody(usrRegExistsMsg)
		return responseBuilder.Build()
	}

	return createUser(user, client, responseBuilder)
}

func createUser(user models.User, client interfaces.DataOrigin, responseBuilder interfaces.ResponseBuilder[events.APIGatewayProxyResponse]) *events.APIGatewayProxyResponse {
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
	return responseBuilder.Build()
}

func userExists(email string, client interfaces.DataOrigin) bool {
	val, _ := client.GetRecord(constants.UsersOrigin, []any{"email", email})

	return val != nil
}

func getClient(ctx *context.Context) interfaces.DataOrigin {
	dbOrigin := utils.GetContextValue[interfaces.DBConnectManager](ctx, constants.CtxKeyDbManager)
	return dbOrigin.GetDataOrigin()
}
