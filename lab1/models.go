package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

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

func migrate() {
	dsn := "host=localhost user=postgres password=568219 dbname=golang port=5432 sslmode=disable TimeZone=Asia/Yekaterinburg"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&Faculty{}, &Specialization{}, &Teacher{}, &Student{}, &Course{}, &Email{})
	if err != nil {
		panic("failed to migrate")
	}
}
