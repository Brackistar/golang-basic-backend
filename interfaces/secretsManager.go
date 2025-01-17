package interfaces

import "github.com/Brackistar/golang-basic-backend/shared/models"

type SecretsManager interface {
	GetSecrets(name string) (models.Secret, error)
}
