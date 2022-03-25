package csv

import (
	"encoding/csv"
	"os"

	models "github.com/BatuhanSerin/books/models/author"
)

func ReadCSV(filename string) (models.AuthorSlice, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	elements, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var books models.AuthorSlice
	for _, line := range elements[1:] {
		books = append(books, models.Author{
			ID:         line[0],
			Name:       line[1],
			Page:       line[2],
			Stock:      line[3],
			Cost:       line[4],
			StockCode:  line[5],
			ISBN:       line[6],
			AuthorID:   line[7],
			AuthorName: line[8],
		})
	}

	return books, nil
}
