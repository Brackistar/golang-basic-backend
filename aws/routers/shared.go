package routers

import (
	"log"

	"github.com/Brackistar/golang-basic-backend/interfaces"
	"github.com/Brackistar/golang-basic-backend/shared/constants"
	"github.com/Brackistar/golang-basic-backend/shared/models"
)

func getUser(email string, client interfaces.DataOrigin) (models.User, error) {
	log.Printf("Retrieving user by email: %s", email)

	var user models.User

	val, err := client.GetRecord(constants.UsersOrigin, "email", email)

	if err != nil {
		log.Printf("Failed to find user by Email: \"%s\"", email)
		log.Print(err.Error())

		return user, err
	}

	log.Printf("User found with email: %s", email)

	user = val.(models.User)

	log.Printf("User found: %+v", user)

	return user, nil
}
