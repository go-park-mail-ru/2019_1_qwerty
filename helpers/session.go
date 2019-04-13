package helpers

import (
	"2019_1_qwerty/models"

	uuid "github.com/satori/uuid"
)

func CreateSession(user string) string {
	sessionID := uuid.NewV4()
	models.Sessions[sessionID.String()] = user
	return sessionID.String()
}

func DestroySession(sessionID string) {
	delete(models.Sessions, sessionID)
}

func ValidateSession(sessionID string) bool {
	_, ok := models.Sessions[sessionID]
	return ok
}

func GetOwner(sessionID string) string {
	res, _ := models.Sessions[sessionID]
	return res
}
