package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Brackistar/golang-basic-backend/aws/configmanager"
	awsconst "github.com/Brackistar/golang-basic-backend/aws/constants"
	"github.com/Brackistar/golang-basic-backend/aws/handlers"
	responsebuilder "github.com/Brackistar/golang-basic-backend/aws/responseBuilder"
	"github.com/Brackistar/golang-basic-backend/aws/secretsmanager"
	"github.com/Brackistar/golang-basic-backend/db"
	"github.com/Brackistar/golang-basic-backend/interfaces"
	"github.com/Brackistar/golang-basic-backend/shared/constants"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
)

var ResponseBuilder interfaces.ResponseBuilder[events.APIGatewayProxyResponse]
var SecretsManager interfaces.SecretsManager
var ConfigManager interfaces.ConfigurationManager[aws.Config]
var DbManager interfaces.DBConnectManager

func main() {
	log.Println("Starting lambda")
	log.Println("Initializing ResponseBuilder")
	ResponseBuilder = responsebuilder.NewAWSResponseBuilder()
	log.Println("Initializing ConfigManager")
	ConfigManager = configmanager.NewAwsConfigManager()
	log.Println("Initializing SecretsManager")
	SecretsManager = secretsmanager.NewAWSSecretsManager(ConfigManager)
	log.Println("Initializing DBManager")
	DbManager = db.NewMongoConnectManager()

	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (event *events.APIGatewayProxyResponse, err error) {

	log.Println("Handling Request")

	// General error handling
	defer func() {
		if r := recover(); r != nil {
			log.Printf(awsconst.FatalErrorMsg, r)

			event = &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       fmt.Sprint(r),
				Headers:    map[string]string{"Content-Type": "application/json"},
			}
			err = fmt.Errorf("%s", r)
		}
	}()

	// Check environment
	if err := validateEnvironment(); err != nil {
		return buildInternalServerErrorResponse(ResponseBuilder, err)
	}

	// Initialize AWS Configuration
	ConfigManager.InitConfig()

	// Get Secrets
	secrets, err := SecretsManager.GetSecrets(os.Getenv(awsconst.SecretName))

	if err != nil {
		return buildInternalServerErrorResponse(ResponseBuilder, err)
	}

	log.Println("Configuring context")
	path := strings.Replace(request.PathParameters["messageboard"], os.Getenv(awsconst.UrlPrefix), "", -1)

	con := ConfigManager.GetContext()
	*con = context.WithValue(*con, constants.CtxKeyPath, path)
	*con = context.WithValue(*con, constants.CtxKeyMethod, request.HTTPMethod)
	*con = context.WithValue(*con, constants.CtxKeyUser, secrets.UserName)
	*con = context.WithValue(*con, constants.CtxKeyPswd, secrets.Password)
	*con = context.WithValue(*con, constants.CtxKeyHost, secrets.Host)
	*con = context.WithValue(*con, constants.CtxKeyDb, secrets.Database)
	*con = context.WithValue(*con, constants.CtxKeyJwt, secrets.Jwt)
	*con = context.WithValue(*con, constants.CtxKeyBdy, request.Body)
	*con = context.WithValue(*con, constants.CtxKeyBckt, os.Getenv(awsconst.Bucket))

	log.Print("Context loaded")

	// Connect with database
	err = DbManager.Connect(*con)

	if err != nil {
		return buildInternalServerErrorResponse(ResponseBuilder, err)
	}

	*con = context.WithValue(*con, constants.CtxKeyDbManager, DbManager)

	return handlers.HandleRequest(con, &request, ResponseBuilder), nil
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

	if err := lookUpEnvironmentVariable(awsconst.UrlPrefix); err != nil {
		return err
	}

	if err := lookUpEnvironmentVariable(awsconst.Bucket); err != nil {
		return err
	}

	if err := lookUpEnvironmentVariable(awsconst.SecretName); err != nil {
		return err
	}

	return nil
}

func lookUpEnvironmentVariable(variableName string) error {
	if _, found := os.LookupEnv(variableName); !found {
		return fmt.Errorf(awsconst.EnvVaNotFoundErrorMsg, variableName)
	}

	return nil
}
