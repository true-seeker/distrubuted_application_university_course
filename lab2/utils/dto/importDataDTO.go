package dto

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ImportDataDTO struct {
	Faculties       []FacultyDTO        `json:"faculties,omitempty"`
	Specializations []SpecializationDTO `json:"specializations,omitempty"`
	Teachers        []TeacherDTO        `json:"teachers,omitempty"`
	Courses         []CourseDTO         `json:"courses,omitempty"`
	Emails          []EmailDTO          `json:"emails,omitempty"`
	Students        []StudentDTO        `json:"students,omitempty"`
}

func importFromDataDTO(dto ImportDataDTO) int {
	jsonData, err := json.Marshal(dto)
	if err != nil {
		panic("Failed to serialize")
	}

	post, err := http.Post("http://localhost:80/excel_import", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic("Failed to post")

	}

	return post.StatusCode
}
