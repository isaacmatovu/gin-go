package config

import (
	"fmt"
	"log"
	"os"
	"quickstart/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB //global variable to hold our database connection


func ConnectDatabase(){

	//load .env file
	err := godotenv.Load()
	if err !=nil{
		log.Println("NO .env file found")
	}
// Build connection string from environment variables
	 dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
    )
	//database connection string
	    // dsn := "host=localhost user=postgres password=Grand@1999 dbname=bookstore port=5432 sslmode=disable"




		//connect to database
		database,err:=gorm.Open(postgres.Open(dsn),&gorm.Config{})
		if err !=nil{
			panic("failed to connect database")
		}
 


			// Drop the join table explicitly FIRST
	err = database.Migrator().DropTable("book_categories")
	if err != nil {
		log.Println("Note: book_categories table may not exist yet:", err)
	}

   // Drop existing tables to recreate with proper constraints
		err = database.Migrator().DropTable(&models.Book{}, &models.Author{},&models.Category{})
		if err != nil {
			log.Fatal("failed to drop tables:", err)
		}

		//automigrate the schema
		err = database.AutoMigrate(&models.Author{},&models.Book{},models.Category{})
		if err !=nil{
			log.Fatal("failed to migrate database:",err)
		}
		fmt.Println("database migration completed successfully")
		DB = database
		fmt.Println("Database connected successfully")
}