package services

import (
	"lab2/utils/services"
)

func SendDataViaQueue(encryptionType string) {
	unnormalizedStudents := services.ReadSqlite("../db.db")
	services.PutUnnormalizedDataToQueue(unnormalizedStudents, encryptionType)
}

func SendDataViaSocket(encryptionType string) {
	unnormalizedStudents := services.ReadSqlite("../db.db")
	services.PutUnnormalizedDataToSocket(unnormalizedStudents, encryptionType)
}
