package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres

var DB *gorm.DB
var err error

type Movie struct {
	gorm.Model
	ID				string		`gorm:"primarykey"`
	Title			string	
	Genre			string		`gorm:"autoCreateTime:false"`
	CreatedAt	time.Time	`gorm:"autoUpdateTime:false"`
}


func DatabaseConnection() {
	host := "localhost"
	port := "5432"
	dbName := "mydb"
	dbUser := "josh"
   password := "airfoil0"
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
       host,
       port,
       dbUser,
       dbName,
       password,
   )

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(Movie{})
	

	if err != nil {
		log.Fatal("Error connecting to database", err)
	}
	fmt.Println("Database connection successful")
}