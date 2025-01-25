package jwt

import (
	"fmt"
	"strings"

	"github.com/Brackistar/golang-basic-backend/shared/models"
	jwt "github.com/golang-jwt/jwt/v5"
)

const (
	bearer     string = "Bearer"
	invFrmtMsg string = "invalid token format"
	invTknMsg  string = "invalid token"
)

var Email string
var UserId string

func HandleToken(token string, jwtSign string) (*models.Claim, bool, string, error) {
	pass := []byte(jwtSign)
	var claims models.Claim

	splitToken := strings.Split(token, bearer)

	if len(splitToken) != 2 {
		return getFailedResponse(invFrmtMsg)
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return pass, nil
	})

	if err != nil {
		model, success, content := getPartialFailResponse()
		return model, success, content, err
	}

	if !tkn.Valid {
		return getFailedResponse(invTknMsg)
	}

	model, success, content := getPartialFailResponse()

	return model, success, content, nil
}

func getFailedResponse(message string) (*models.Claim, bool, string, error) {
	model, success, content := getPartialFailResponse()
	return model, success, content, fmt.Errorf(message)
}

func getPartialFailResponse() (*models.Claim, bool, string) {
	return &models.Claim{}, false, string("")
}
