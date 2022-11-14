package services

import "lab2/utils/services"

func AcceptDataFromQueue() {
	services.GetUnnormalizedDataFromQueue()
}

func AcceptDataFromSocket() {
	services.GetUnnormalizedDataFromSocket()
}
