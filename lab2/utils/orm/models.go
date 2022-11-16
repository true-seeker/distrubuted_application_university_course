package orm

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"lab2/utils/config"
	"time"
)

var PostgresConnectionString = fmt.Sprintf("host=localhost "+
	"user=%s "+
	"password=%s "+
	"dbname=%s "+
	"port=%s "+
	"sslmode=disable TimeZone=Asia/Yekaterinburg",
	config.GetProperty("DataBase", "user"),
	config.GetProperty("DataBase", "password"),
	config.GetProperty("DataBase", "dbname"),
	config.GetProperty("DataBase", "port"))

type Faculty struct {
	gorm.Model
	Title string
}

type Specialization struct {
	gorm.Model
	FacultyId uint
	Faculty   Faculty `gorm:"foreignKey:FacultyId;references:ID"`
	Title     string
}

type Teacher struct {
	gorm.Model
	UnnormalizedId string
	Name           string
}

type Course struct {
	gorm.Model
	Title     string
	FacultyId uint
	Faculty   Faculty `gorm:"foreignKey:FacultyId;references:ID"`
	TeacherId uint
	Teacher   Teacher `gorm:"foreignKey:TeacherId;references:ID"`
}

type Email struct {
	gorm.Model
	Mail      string
	StudentId uint
}

type Student struct {
	gorm.Model
	Name             string
	BirthDate        time.Time
	SpecializationId uint
	Specialization   Specialization `gorm:"foreignKey:SpecializationId;references:ID"`
	Courses          []Course       `gorm:"many2many:student_courses;"`
	Emails           []Email
	UnnormalizedId   string
}

func Migrate() {
	db, err := gorm.Open(postgres.Open(PostgresConnectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&Faculty{}, &Specialization{}, &Teacher{}, &Student{}, &Course{}, &Email{})
	if err != nil {
		panic("failed to migrate")
	}
}
