package dto

import (
	"lab3/utils/orm"
	"time"
)

type StudentDTO struct {
	Id             uint
	Name           string
	BirthDate      time.Time
	Specialization SpecializationDTO
	Courses        []CourseDTO
	Emails         []EmailDTO
}

func MapStudentsDTO(students []orm.Student) (dtos []StudentDTO) {
	for i := 0; i < len(students); i++ {
		dtos = append(dtos, MapStudentToDTO(students[i]))
	}
	return
}

func MapStudentToDTO(student orm.Student) (dto StudentDTO) {
	dto = StudentDTO{
		Id:             student.ID,
		Name:           student.Name,
		BirthDate:      student.BirthDate,
		Specialization: MapSpecializationDTO(student.Specialization),
		Courses:        MapCoursesDTO(student.Courses),
		Emails:         MapEmailsDTO(student.Emails),
	}
	return
}
