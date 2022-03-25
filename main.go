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
	bookRepo.InsertData(books)

	fmt.Println(bookRepo.FindAll())

	fmt.Println(bookRepo.GetByID(3))

	fmt.Println(bookRepo.FindByName("White Fang"))

	fmt.Println(bookRepo.GetBooksWithAuthor("Jack London"))

	//bookRepo.DeleteById(2)

	fmt.Println(bookRepo.FindAll())

}
