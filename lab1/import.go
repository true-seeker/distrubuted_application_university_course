package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"lab1/dto"
	"net/http"
)

type importDTO struct {
	Faculties       []dto.FacultyDTO        `json:"faculties,omitempty"`
	Specializations []dto.SpecializationDTO `json:"specializations,omitempty"`
	Teachers        []dto.TeacherDTO        `json:"teachers,omitempty"`
	Courses         []dto.CourseDTO         `json:"courses,omitempty"`
	Emails          []dto.EmailDTO          `json:"emails,omitempty"`
	Students        []dto.StudentDTO        `json:"students,omitempty"`
}

func importFromDTO(dto importDTO) int {
	jsonData, err := json.Marshal(dto)
	if err != nil {
		panic("Failed to serialize")
	}

	post, err := http.Post("http://localhost:80/excel_import", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
		panic("Failed to post")

	}

	return post.StatusCode
}
