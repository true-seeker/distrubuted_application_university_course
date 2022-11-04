package services

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"lab2/utils/dto"
	"lab2/utils/orm"
	"strings"
	"time"
)

func NormalizeStudents(students []dto.UnnormalizedStudent) {
	db, err := gorm.Open(postgres.Open(orm.PostgresConnectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Delete(&orm.Faculty{}, "deleted_at is null")
	db.Delete(&orm.Specialization{}, "deleted_at is null")
	db.Delete(&orm.Teacher{}, "deleted_at is null")
	db.Delete(&orm.Course{}, "deleted_at is null")
	db.Delete(&orm.Student{}, "deleted_at is null")

	for i := 0; i < len(students); i++ {
		normalizeStudent(students[i], db)
	}
}

func normalizeStudent(student dto.UnnormalizedStudent, db *gorm.DB) {

	var faculty orm.Faculty
	db.First(&faculty, "Title = ?", student.Faculty)
	if faculty.ID == 0 {
		db.Create(&orm.Faculty{
			Title: student.Faculty,
		})
	}
	db.First(&faculty, "Title = ?", student.Faculty)

	var specialization orm.Specialization
	db.First(&specialization, "Title = ?", student.Specialization)
	if specialization.ID == 0 {
		db.Create(&orm.Specialization{
			FacultyId: faculty.ID,
			Title:     student.Specialization,
		})
	}
	db.First(&specialization, "Title = ?", student.Specialization)

	teachers := strings.Split(student.Teachers, "|")
	var ORMTearchers []orm.Teacher
	for i := 0; i < len(teachers); i++ {
		var teacher orm.Teacher
		teacherData := strings.Split(teachers[i], ",")

		db.First(&teacher, "unnormalized_id = ?", teacherData[0])
		if teacher.ID == 0 {
			db.Create(&orm.Teacher{
				UnnormalizedId: teacherData[0],
				Name:           teacherData[1],
			})
		}
		db.First(&teacher, "unnormalized_id = ?", teacherData[0])
		ORMTearchers = append(ORMTearchers, teacher)
	}

	courses := strings.Split(student.Courses, "|")
	var ORMcourses []orm.Course
	for i := 0; i < len(courses); i++ {
		var course orm.Course
		db.First(&course, "Title = ?", courses[i])
		if course.ID == 0 {
			db.Create(&orm.Course{
				Title:     courses[i],
				FacultyId: faculty.ID,
				TeacherId: ORMTearchers[i].ID,
			})
			db.First(&course, "Title = ?", courses[i])
		}
		ORMcourses = append(ORMcourses, course)
	}

	studentData := strings.Split(student.Name, ",")
	var ORMStudent orm.Student
	db.First(&ORMStudent, "unnormalized_id = ?", studentData[0])
	if ORMStudent.ID == 0 {
		parsedDate, _ := time.Parse("01-02-2006", student.BirthDate)
		db.Create(&orm.Student{
			Name:             studentData[1],
			BirthDate:        parsedDate,
			SpecializationId: specialization.ID,
			UnnormalizedId:   studentData[0],
			Courses:          ORMcourses,
		})
	}
	db.First(&ORMStudent, "unnormalized_id = ?", studentData[0])

	emails := strings.Split(student.Emails, "|")
	var ORMemails []orm.Email
	for i := 0; i < len(emails); i++ {
		var email orm.Email
		db.First(&email, "Mail = ?", emails[i])
		if email.ID == 0 {
			db.Create(&orm.Email{
				Mail:      emails[i],
				StudentId: ORMStudent.ID,
			})
			db.First(&email, "Mail = ?", emails[i])
		}
		ORMemails = append(ORMemails, email)
	}
	ORMStudent.Emails = ORMemails
	db.Save(&ORMStudent)
}

func SendNormalizedDataToImport() {
	db, err := gorm.Open(postgres.Open(orm.PostgresConnectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var faculties []orm.Faculty
	db.Find(&faculties)
	facultiesDTO := dto.MapFacultiesDTO(faculties)

	var specializations []orm.Specialization
	db.Preload("Faculty").
		Find(&specializations)
	specializationsDTO := dto.MapSpecializationsDTO(specializations)

	var teachers []orm.Teacher
	db.Find(&teachers)
	teachersDTO := dto.MapTeachersDTO(teachers)

	var courses []orm.Course
	db.Preload("Faculty").
		Preload("Teacher").
		Find(&courses)
	coursesDTO := dto.MapCoursesDTO(courses)

	var emails []orm.Email
	db.Find(&emails)
	emailsDTO := dto.MapEmailsDTO(emails)

	var students []orm.Student
	db.Preload("Courses").
		Preload("Courses.Teacher").
		Preload("Courses.Faculty").
		Preload("Specialization").
		Preload("Specialization.Faculty").
		Preload("Emails").
		Find(&students)
	studentsDTO := dto.MapStudentsDTO(students)

	importDataDTO := dto.ImportDataDTO{
		Faculties:       facultiesDTO,
		Specializations: specializationsDTO,
		Teachers:        teachersDTO,
		Courses:         coursesDTO,
		Emails:          emailsDTO,
		Students:        studentsDTO,
	}

	var fields = MakeFieldsArray()

	MakeImportRequest(dto.ImportDTO{
		Data:   importDataDTO,
		Fields: fields,
	})
}
