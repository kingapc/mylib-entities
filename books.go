package entities

import (
	conn "github.com/rpinedafocus/mylib-dbconn"
)

type Books struct {
	NAME        string `json:"name"`
	AKA         string `json:"aka"`
	TOTAL       int    `json:"total"`
	PUBLIS_DATE string `json:"publish_date"`
	AUTHOR      string `json:"author"`
	GENRE       string `json:"genre"`
}

func GetBooks() []Books {

	var book Books
	var books []Books
	db, errc := conn.GetConnection()

	if errc {
		return books
	} else {
		rows, err := db.Query("SELECT a.name_book NAME, a.second_name AKA, a.total TOTAL, a.publish_date PUBLIS_DATE, b.name AUTHOR, c.name GENRE FROM university.books a INNER JOIN university.authors b ON a.author_id = b.author_id INNER JOIN university.genres c ON a.genre_id = c.genre_id")

		if err != nil {
			return books
		} else {
			for rows.Next() {
				err := rows.Scan(&book.NAME, &book.AKA, &book.TOTAL, &book.PUBLIS_DATE, &book.AUTHOR, &book.GENRE)
				if err != nil {
					return books
				}

				books = append(books, book)
			}

			return books
		}
	}
}
