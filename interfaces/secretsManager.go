package interfaces

import "github.com/Brackistar/golang-basic-backend/models"

type SecretsManager interface {
	GetSecrets(name string) (models.Secret, error)
}
