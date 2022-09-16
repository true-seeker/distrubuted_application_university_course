package dto

import "lab1/orm"

type FacultyDTO struct {
	Id    uint
	Title string
}

func MapFacultiesDTO(faculties []orm.Faculty) (dtos []FacultyDTO) {
	for i := 0; i < len(faculties); i++ {
		dtos = append(dtos, MapFacultyToDTO(faculties[i]))
	}
	return
}

func MapFacultyToDTO(faculty orm.Faculty) (dto FacultyDTO) {
	dto = FacultyDTO{
		Title: faculty.Title,
		Id:    faculty.ID,
	}
	return
}
