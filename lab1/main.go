package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
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

func main() {
	unnormalizedStudents := readSqlite("./db.db")
	fmt.Println(unnormalizedStudents)
}
