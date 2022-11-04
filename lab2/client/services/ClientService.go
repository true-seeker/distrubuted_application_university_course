package services

import (
	"fmt"
	"lab2/utils/services"
)

func SendDataViaQueue(encryptionType string) {
	unnormalizedStudents := services.ReadSqlite("db.db")
	fmt.Println(unnormalizedStudents)
}
