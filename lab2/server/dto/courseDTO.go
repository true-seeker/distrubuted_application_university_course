package dto

import "server/orm"

type CourseDTO struct {
	Id      uint
	Title   string
	Faculty FacultyDTO
	Teacher TeacherDTO
}

func MapCoursesDTO(courses []orm.Course) (dtos []CourseDTO) {
	for i := 0; i < len(courses); i++ {
		dtos = append(dtos, MapCourseToDTO(courses[i]))
	}
	return
}

func MapCourseToDTO(course orm.Course) (dto CourseDTO) {
	dto = CourseDTO{
		Id:      course.ID,
		Title:   course.Title,
		Faculty: MapFacultyToDTO(course.Faculty),
		Teacher: MapTeacherToDTO(course.Teacher),
	}
	return
}
