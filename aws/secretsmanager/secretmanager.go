package secretsmanager

import (
	"encoding/json"
	"log"

	"github.com/Brackistar/golang-basic-backend/interfaces"
	"github.com/Brackistar/golang-basic-backend/shared/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type AWSSecretsManager struct {
	ConfigManager interfaces.ConfigurationManager[aws.Config]
}

func NewAWSSecretsManager(configManager interfaces.ConfigurationManager[aws.Config]) *AWSSecretsManager {
	return &AWSSecretsManager{
		ConfigManager: configManager,
	}
}

func (i *AWSSecretsManager) GetSecrets(name string) (models.Secret, error) {
	var result models.Secret

	log.Printf("Looking for information on secret: \"%s\"", name)

	log.Println("Loading AWS Configuration")
	svc := secretsmanager.NewFromConfig(i.ConfigManager.GetConfig())

	log.Println("Downloading secrets")
	pass, err := svc.GetSecretValue(*i.ConfigManager.GetContext(), &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(name),
	})

	if err != nil {
		log.Print(err)
		return result, err
	}

	json.Unmarshal([]byte(*pass.SecretString), &result)

	log.Print("Secrets information gathered")

	return result, nil
}
