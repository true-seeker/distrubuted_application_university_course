package services

import (
	"bytes"
	"encoding/json"
	"lab1/dto"
	"net/http"
)

func MakeImportRequest(dto dto.ImportDTO) int {
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

func MakeFieldsArray() (dtos []dto.SheetFieldDTO) {

	facultyEntityFields := []dto.EntityFieldDTO{{
		FieldName: "Id",
		Title:     "Id",
	}, {
		FieldName: "Title",
		Title:     "Наименование",
	}}
	dtos = append(dtos, dto.SheetFieldDTO{
		FieldName:    "faculties",
		Title:        "Факультеты",
		EntityFields: facultyEntityFields,
	})

	specializationEntityFields := []dto.EntityFieldDTO{{
		FieldName: "Id",
		Title:     "Id",
	}, {
		FieldName: "Title",
		Title:     "Наименование",
	}, {
		FieldName: "Faculty.Id",
		Title:     "Id факультета",
	}, {
		FieldName: "Faculty.Title",
		Title:     "Наименование факультета",
	}}
	dtos = append(dtos, dto.SheetFieldDTO{
		FieldName:    "specializations",
		Title:        "Специализации",
		EntityFields: specializationEntityFields,
	})

	teacherEntityFields := []dto.EntityFieldDTO{{
		FieldName: "Id",
		Title:     "Id",
	}, {
		FieldName: "Name",
		Title:     "Имя",
	}}
	dtos = append(dtos, dto.SheetFieldDTO{
		FieldName:    "teachers",
		Title:        "Преподаватели",
		EntityFields: teacherEntityFields,
	})

	courseEntityFields := []dto.EntityFieldDTO{{
		FieldName: "Id",
		Title:     "Id",
	}, {
		FieldName: "Title",
		Title:     "Наименование",
	}, {
		FieldName: "Faculty.Id",
		Title:     "Id факультета",
	}, {
		FieldName: "Faculty.Title",
		Title:     "Наименование факультета",
	}}
	dtos = append(dtos, dto.SheetFieldDTO{
		FieldName:    "courses",
		Title:        "Предметы",
		EntityFields: courseEntityFields,
	})

	emailEntityFields := []dto.EntityFieldDTO{{
		FieldName: "Id",
		Title:     "Id",
	}, {
		FieldName: "Mail",
		Title:     "Почта",
	}, {
		FieldName: "StudentId",
		Title:     "Id студента",
	}}
	dtos = append(dtos, dto.SheetFieldDTO{
		FieldName:    "emails",
		Title:        "Электронные почты",
		EntityFields: emailEntityFields,
	})

	studentEntityFields := []dto.EntityFieldDTO{{
		FieldName: "Id",
		Title:     "Id",
	}, {
		FieldName: "Name",
		Title:     "Имя",
	}, {
		FieldName: "BirthDate",
		Title:     "Дата рождения",
	}, {
		FieldName: "Specialization.Id",
		Title:     "Id специализации",
	}, {
		FieldName: "Specialization.Title",
		Title:     "Наименование специализации",
	}, {
		FieldName: "Specialization.Faculty.Id",
		Title:     "Id факультета",
	}, {
		FieldName: "Specialization.Faculty.Title",
		Title:     "Наименование факультета",
	}}
	dtos = append(dtos, dto.SheetFieldDTO{
		FieldName:    "students",
		Title:        "Студенты",
		EntityFields: studentEntityFields,
	})

	return
}
