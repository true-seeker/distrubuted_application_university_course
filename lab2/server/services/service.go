package services

func AcceptDataFromQueue(encryptionType string) {
	GetUnnormalizedDataFromQueue(encryptionType)
}

func AcceptDataFromSocket(encryptionType string) {
	GetUnnormalizedDataFromSocket(encryptionType)
}
