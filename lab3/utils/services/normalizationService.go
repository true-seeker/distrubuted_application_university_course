package services

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"lab3/utils/dto"
	"lab3/utils/orm"
	"strings"
	"time"
)

func NormalizeStudent(student dto.UnnormalizedStudent) {
	db, _ := gorm.Open(postgres.Open(orm.PostgresConnectionString), &gorm.Config{})

	var faculty orm.Faculty
	db.First(&faculty, "Title = ?", student.Faculty)
	db.Create(&orm.Faculty{
		Title: student.Faculty,
	})
	db.First(&faculty, "Title = ?", student.Faculty)

	var specialization orm.Specialization
	db.First(&specialization, "Title = ?", student.Specialization)
	db.Create(&orm.Specialization{
		FacultyId: faculty.ID,
		Title:     student.Specialization,
	})
	db.First(&specialization, "Title = ?", student.Specialization)

	teachers := strings.Split(student.Teachers, "|")
	var ORMTearchers []orm.Teacher
	for i := 0; i < len(teachers); i++ {
		var teacher orm.Teacher
		teacherData := strings.Split(teachers[i], ",")

		db.First(&teacher, "unnormalized_id = ?", teacherData[0])
		db.Create(&orm.Teacher{
			UnnormalizedId: teacherData[0],
			Name:           teacherData[1],
		})
		db.First(&teacher, "unnormalized_id = ?", teacherData[0])
		ORMTearchers = append(ORMTearchers, teacher)
	}

	courses := strings.Split(student.Courses, "|")
	var ORMcourses []orm.Course
	for i := 0; i < len(courses); i++ {
		var course orm.Course
		db.First(&course, "Title = ?", courses[i])
		db.Create(&orm.Course{
			Title:     courses[i],
			FacultyId: faculty.ID,
			TeacherId: ORMTearchers[i].ID,
		})
		db.First(&course, "Title = ?", courses[i])
		ORMcourses = append(ORMcourses, course)
	}

	studentData := strings.Split(student.Name, ",")
	var ORMStudent orm.Student
	db.First(&ORMStudent, "unnormalized_id = ?", studentData[0])
	parsedDate, _ := time.Parse("01-02-2006", student.BirthDate)
	db.Create(&orm.Student{
		Name:             studentData[1],
		BirthDate:        parsedDate,
		SpecializationId: specialization.ID,
		UnnormalizedId:   studentData[0],
		Courses:          ORMcourses,
	})
	db.First(&ORMStudent, "unnormalized_id = ?", studentData[0])

	emails := strings.Split(student.Emails, "|")
	var ORMemails []orm.Email
	for i := 0; i < len(emails); i++ {
		var email orm.Email
		db.First(&email, "Mail = ?", emails[i])
		db.Create(&orm.Email{
			Mail:      emails[i],
			StudentId: ORMStudent.ID,
		})
		db.First(&email, "Mail = ?", emails[i])
		ORMemails = append(ORMemails, email)
	}
	ORMStudent.Emails = ORMemails
	db.Save(&ORMStudent)
}
