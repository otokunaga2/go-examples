package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jinzhu/gorm"
	"time"
)

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type TextLog struct {
	ID        uint `gorm:"primary_key"`
	Comment   string
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=test password=secret sslmode=disable")
	if err != nil {
		fmt.Println("Fail we miss something", err)
	}
	db.CreateTable(&TextLog{})

	defer db.Close()
}
