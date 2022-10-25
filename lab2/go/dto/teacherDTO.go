package dto

import "lab2/orm"

type TeacherDTO struct {
	Id   uint
	Name string
}

func MapTeachersDTO(teachers []orm.Teacher) (dtos []TeacherDTO) {
	for i := 0; i < len(teachers); i++ {
		dtos = append(dtos, MapTeacherToDTO(teachers[i]))
	}
	return
}

func MapTeacherToDTO(teacher orm.Teacher) (dto TeacherDTO) {
	dto = TeacherDTO{
		Id:   teacher.ID,
		Name: teacher.Name,
	}
	return
}
