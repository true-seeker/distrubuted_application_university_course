package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
	"time"
)

type unnormalizedStudent struct {
	id             int
	name           string
	emails         string
	courses        string
	birthDate      string
	teachers       string
	faculty        string
	specialization string
}

func readSqlite(dbName string) []unnormalizedStudent {
	var uss []unnormalizedStudent
	db, _ := sql.Open("sqlite3", dbName)
	defer db.Close()

	rows, _ := db.Query("SELECT id, name, emails, courses, birth_date, teachers, faculty, specialization FROM students")

	defer rows.Close()

	for rows.Next() {
		us := unnormalizedStudent{}

		err := rows.Scan(&us.id, &us.name, &us.emails, &us.courses, &us.birthDate, &us.teachers, &us.faculty, &us.specialization)
		if err != nil {
			continue
		}

		uss = append(uss, us)
	}
	return uss
}

func normalizeStudents(students []unnormalizedStudent) {
	dsn := "host=localhost user=postgres password=568219 dbname=golang port=5432 sslmode=disable TimeZone=Asia/Yekaterinburg"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.Delete(&Faculty{}, "deleted_at is null")
	db.Delete(&Specialization{}, "deleted_at is null")
	db.Delete(&Teacher{}, "deleted_at is null")
	db.Delete(&Course{}, "deleted_at is null")
	db.Delete(&Student{}, "deleted_at is null")

	if err != nil {
		panic("failed to connect database")
	}
	for i := 0; i < len(students); i++ {
		fmt.Println(students[i])
		normalizeStudent(students[i], db)
	}
}

func normalizeStudent(student unnormalizedStudent, db *gorm.DB) {

	var faculty Faculty
	db.First(&faculty, "Title = ?", student.faculty)
	if faculty.ID == 0 {
		db.Create(&Faculty{
			Title: student.faculty,
		})
	}
	db.First(&faculty, "Title = ?", student.faculty)

	var specialization Specialization
	db.First(&specialization, "Title = ?", student.specialization)
	if specialization.ID == 0 {
		db.Create(&Specialization{
			FacultyId: faculty.ID,
			Title:     student.specialization,
		})
	}
	db.First(&specialization, "Title = ?", student.specialization)

	teachers := strings.Split(student.teachers, "|")
	var ORMTearchers []Teacher
	for i := 0; i < len(teachers); i++ {
		var teacher Teacher
		teacherData := strings.Split(teachers[i], ",")

		db.First(&teacher, "unnormalized_id = ?", teacherData[0])
		if teacher.ID == 0 {
			db.Create(&Teacher{
				UnnormalizedId: teacherData[0],
				Name:           teacherData[1],
			})
		}
		db.First(&teacher, "unnormalized_id = ?", teacherData[0])
		ORMTearchers = append(ORMTearchers, teacher)
	}

	courses := strings.Split(student.courses, "|")
	for i := 0; i < len(courses); i++ {
		var course Course
		db.First(&course, "Title = ?", courses[i])
		if course.ID == 0 {
			db.Create(&Course{
				Title:     courses[i],
				FacultyId: faculty.ID,
				TeacherId: ORMTearchers[i].ID,
			})
		}
		db.First(&course, "Title = ?", courses[i])
	}

	studentData := strings.Split(student.name, ",")
	var ORMStudent Student
	db.First(&ORMStudent, "unnormalized_id = ?", studentData[0])
	if ORMStudent.ID == 0 {
		parsedDate, _ := time.Parse("31-12-2022", student.birthDate)
		db.Create(&Student{
			Name:             studentData[1],
			BirthDate:        parsedDate,
			SpecializationId: specialization.ID,
			UnnormalizedId:   studentData[0],
		})
	}
	db.First(&ORMStudent, "unnormalized_id = ?", studentData[0])

	emails := strings.Split(student.emails, "|")
	for i := 0; i < len(emails); i++ {
		var email Email
		db.First(&email, "Mail = ?", emails[i])
		if email.ID == 0 {
			db.Create(&Email{
				Mail:      emails[i],
				StudentId: ORMStudent.ID,
			})
		}
	}
}

func main() {
	migrate()
	unnormalizedStudents := readSqlite("./db.db")

	normalizeStudents(unnormalizedStudents)
}
