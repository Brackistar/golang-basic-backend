package aws

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Brackistar/golang-basic-backend/aws/handlers"
	"github.com/Brackistar/golang-basic-backend/interfaces"
	"github.com/Brackistar/golang-basic-backend/shared/constants"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type AWSRequestHandler struct {
	ResponseBuilder interfaces.ResponseBuilder[events.APIGatewayProxyResponse]
	SecretsManager  interfaces.SecretsManager
	ConfigManager   interfaces.ConfigurationManager[aws.Config]
	DbManager       interfaces.DBConnectManager
}

func (i *AWSRequestHandler) HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	// General error handling
	defer func() (*events.APIGatewayProxyResponse, error) {
		i.ResponseBuilder.Clear()

		if r := recover(); r != nil {
			log.Printf(fatalErrorMsg, r)

			i.ResponseBuilder.SetStatusCode(http.StatusInternalServerError)
			i.ResponseBuilder.SetBody(genFatalErrorMsg)
			i.ResponseBuilder.AddHeader(getJsonContentHeader())
		}

		return i.ResponseBuilder.Build(), nil
	}()

	// Check environment
	if err := validateEnvironment(); err != nil {
		return buildInternalServerErrorResponse(i.ResponseBuilder, err)
	}

	// Initialize AWS Configuration
	i.ConfigManager.InitConfig()

	// Get Secrets
	secrets, err := i.SecretsManager.GetSecrets(os.Getenv(secretName))

	if err != nil {
		return buildInternalServerErrorResponse(i.ResponseBuilder, err)
	}

	path := strings.Replace(request.PathParameters["messageboard"], os.Getenv(urlPrefix), "", -1)

	con := i.ConfigManager.GetContext()
	*con = context.WithValue(*con, constants.CtxKeyPath, path)
	*con = context.WithValue(*con, constants.CtxKeyMethod, request.HTTPMethod)
	*con = context.WithValue(*con, constants.CtxKeyUser, secrets.UserName)
	*con = context.WithValue(*con, constants.CtxKeyPswd, secrets.Password)
	*con = context.WithValue(*con, constants.CtxKeyHost, secrets.Host)
	*con = context.WithValue(*con, constants.CtxKeyDb, secrets.Database)
	*con = context.WithValue(*con, constants.CtxKeyJwt, secrets.Jwt)
	*con = context.WithValue(*con, constants.CtxKeyBdy, request.Body)
	*con = context.WithValue(*con, constants.CtxKeyBckt, os.Getenv(bucket))

	// Connect with database
	err = i.DbManager.Connect(*con)

	if err != nil {
		return buildInternalServerErrorResponse(i.ResponseBuilder, err)
	}

	*con = context.WithValue(*con, constants.CtxKeyDbManager, i.DbManager)

	return handlers.HandleRequest(con, &request, i.ResponseBuilder), nil
}

// Builds an *events.APIGatewayProxyResponse response using the responce builder provided, with status code 500 and with headers for json content type
func buildInternalServerErrorResponse(builder interfaces.ResponseBuilder[events.APIGatewayProxyResponse], err error) (*events.APIGatewayProxyResponse, error) {
	builder.SetStatusCode(http.StatusInternalServerError)
	builder.SetBody(err.Error())
	builder.AddHeader(getJsonContentHeader())

	return builder.Build(), err
}

func getJsonContentHeader() (string, string) {
	return "Content-Type", "application/json"
}

func validateEnvironment() error {

	if err := lookUpEnvironmentVariable(urlPrefix); err != nil {
		return err
	}

	if err := lookUpEnvironmentVariable(bucket); err != nil {
		return err
	}

	if err := lookUpEnvironmentVariable(secretName); err != nil {
		return err
	}

	return nil
}

func lookUpEnvironmentVariable(variableName string) error {
	if _, found := os.LookupEnv(variableName); !found {
		return fmt.Errorf(envVaNotFoundErrorMsg, variableName)
	}

	return nil
}
