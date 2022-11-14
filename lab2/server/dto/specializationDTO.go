package dto

import "server/orm"

type SpecializationDTO struct {
	Id      uint
	Faculty FacultyDTO
	Title   string
}

func MapSpecializationsDTO(specializations []orm.Specialization) (dtos []SpecializationDTO) {
	for i := 0; i < len(specializations); i++ {
		dtos = append(dtos, MapSpecializationDTO(specializations[i]))
	}
	return
}

func MapSpecializationDTO(specialization orm.Specialization) (dto SpecializationDTO) {
	dto = SpecializationDTO{
		Id:      specialization.ID,
		Faculty: MapFacultyToDTO(specialization.Faculty),
		Title:   specialization.Title,
	}

	return
}
