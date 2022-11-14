package services

func SendDataViaQueue(encryptionType string) {
	unnormalizedStudents := ReadSqlite("db.db")
	PutUnnormalizedDataToQueue(unnormalizedStudents, encryptionType)
}

func SendDataViaSocket(encryptionType string) {
	unnormalizedStudents := ReadSqlite("db.db")
	PutUnnormalizedDataToSocket(unnormalizedStudents, encryptionType)
}
