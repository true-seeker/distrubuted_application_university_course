package main

import (
	_ "github.com/mattn/go-sqlite3"
	"lab2/orm"
	"lab2/services"
)

func main() {
	orm.Migrate()
	unnormalizedStudents := services.ReadSqlite("./db.db")

	services.NormalizeStudents(unnormalizedStudents)

	services.SendNormalizedDataToImport()
}
