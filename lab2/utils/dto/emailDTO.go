package dto

import "lab2/utils/orm"

type EmailDTO struct {
	Id        uint
	Mail      string
	StudentId uint
}

func MapEmailsDTO(emails []orm.Email) (dtos []EmailDTO) {
	for i := 0; i < len(emails); i++ {
		dtos = append(dtos, MapEmailToDTO(emails[i]))
	}
	return
}

func MapEmailToDTO(email orm.Email) (dto EmailDTO) {
	dto = EmailDTO{
		Id:        email.ID,
		Mail:      email.Mail,
		StudentId: email.StudentId,
	}
	return
}
