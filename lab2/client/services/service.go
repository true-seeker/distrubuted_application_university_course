package services

import (
	"lab2/utils/services"
)

func SendDataViaQueue() {
	unnormalizedStudents := services.ReadSqlite("../db.db")
	services.PutUnnormalizedDataToQueue(unnormalizedStudents)
}

func SendDataViaSocket() {
	unnormalizedStudents := services.ReadSqlite("../db.db")
	services.PutUnnormalizedDataToSocket(unnormalizedStudents)
}
