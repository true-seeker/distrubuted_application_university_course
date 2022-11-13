package services

import "lab2/utils/services"

func AcceptDataFromQueue(encryptionType string) {
	services.GetUnnormalizedDataFromQueue()
}

func AcceptDataFromSocket(encryptionType string) {
	services.GetUnnormalizedDataFromSocket()
}
