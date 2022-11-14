package services

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"server/dto"
)

func ReadSqlite(dbName string) []dto.UnnormalizedStudent {
	var uss []dto.UnnormalizedStudent
	db, _ := sql.Open("sqlite3", dbName)
	defer db.Close()

	rows, _ := db.Query("SELECT id, name, emails, courses, birth_date, teachers, faculty, specialization FROM students")

	defer rows.Close()

	for rows.Next() {
		us := dto.UnnormalizedStudent{}

		err := rows.Scan(&us.Id, &us.Name, &us.Emails, &us.Courses, &us.BirthDate, &us.Teachers, &us.Faculty, &us.Specialization)
		if err != nil {
			continue
		}

		uss = append(uss, us)
	}
	return uss
}
