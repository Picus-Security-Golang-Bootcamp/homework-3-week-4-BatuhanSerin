package main

import (
	"fmt"
	"log"

	csv_utils "github.com/BatuhanSerin/books/csv"
	postgres "github.com/BatuhanSerin/books/db"

	"github.com/BatuhanSerin/books/models/author"
	//"github.com/BatuhanSerin/books/models/book"
	"github.com/joho/godotenv"
)

func main() {

	books, err := csv_utils.ReadCSV("book-csv.csv")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(books); i++ {
		fmt.Println(books[i])
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init:", err)
	}
	log.Println("Postgres connected")

	bookRepo := author.NewBookRepository(db)
	bookRepo.Migrations()
	bookRepo.InsertSampleData(books)

	fmt.Println(bookRepo.FindAll())

}
