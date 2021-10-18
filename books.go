package entities

import (
	conn "github.com/rpinedafocus/mylib-dbconn"
)

type Books struct {
	BOOK_ID     int    `json:"id"`
	NAME        string `json:"name"`
	AKA         string `json:"aka"`
	TOTAL       int    `json:"total"`
	PUBLIS_DATE string `json:"publish_date"`
	AUTHOR      string `json:"author"`
	GENRE       string `json:"genre"`
}

type MyBook struct {
	ID           int    `json:"id"`
	NAME_BOOK    string `json:"name_book"`
	STATUS       string `json:"status"`
	PUBLISH_DATE string `json:"publish_date"`
	AUTHOR       string `json:"author"`
	GENRE        string `json:"gere"`
}

func GetBooks() []Books {

	var book Books
	var books []Books
	db, errc := conn.GetConnection()

	if errc {
		return books
	} else {
		rows, err := db.Query("SELECT a.book_id BOOK_ID, TRIM(a.name_book) NAME, COALESCE(a.second_name,'') AKA, a.total TOTAL, a.publish_date PUBLIS_DATE, TRIM(b.name) AUTHOR, TRIM(c.name) GENRE " +
			" FROM university.books a INNER JOIN university.authors b ON a.author_id = b.author_id INNER JOIN university.genres c ON a.genre_id = c.genre_id")

		if err != nil {
			panic(err)
		} else {
			for rows.Next() {
				err := rows.Scan(&book.BOOK_ID, &book.NAME, &book.AKA, &book.TOTAL, &book.PUBLIS_DATE, &book.AUTHOR, &book.GENRE)
				if err != nil {
					panic(err)
				}

				books = append(books, book)
			}

			return books
		}
	}
}

func GetBook(id int) (Books, bool) {

	var book Books
	db, errc := conn.GetConnection()

	if errc {
		return book, false
	} else {
		stmt, err := db.Prepare("SELECT a.book_id BOOK_ID,TRIM(a.name_book) NAME, COALESCE(a.second_name,'') AKA, a.total TOTAL, a.publish_date PUBLIS_DATE, TRIM(b.name) AUTHOR, TRIM(c.name) GENRE " +
			" FROM university.books a INNER JOIN university.authors b ON a.author_id = b.author_id INNER JOIN university.genres c ON a.genre_id = c.genre_id " +
			" WHERE a.book_id = $1")

		if err != nil {
			return book, false
		} else {
			err = stmt.QueryRow(id).Scan(&book.BOOK_ID, &book.NAME, &book.AKA, &book.TOTAL, &book.PUBLIS_DATE, &book.AUTHOR, &book.GENRE)

			if err != nil {
				return book, false
			}

			return book, true
		}
	}
}

func GetMyReservesRents(user string) (MyBook, bool) {

	var book MyBook
	db, errc := conn.GetConnection()

	if errc {
		return book, false
	} else {
		stmt, err := db.Prepare("SELECT a.book_id ID, b.name_book NAME_BOOK, CASE WHEN a.reserved_date IS NULL THEN 'RENTED' ELSE 'RESERVED' END STATUS, b.publish_date PUBLISH_DATE, " +
			" c.name AUTHOR, d.name GENRE FROM university.status a INNER JOIN university.books b ON a.book_id = b.book_id INNER JOIN university.authors c ON c.author_id = b.author_id " +
			" INNER JOIN university.genres d ON d.genre_id = b.genre_id WHERE a.rented_reserved_by = $1")

		if err != nil {
			return book, false
		} else {
			err = stmt.QueryRow(user).Scan(&book.ID, &book.NAME_BOOK, &book.STATUS, &book.PUBLISH_DATE, &book.AUTHOR, &book.GENRE)

			if err != nil {
				return book, false
			}

			return book, true
		}
	}
}

/*
func main() {

	var test = GetBooks()

	// j, err := json.Marshal(test)

	// if err != nil {
	// 	fmt.Printf("Error: %s", err.Error())
	// } else {
	// 	fmt.Println(string(j))
	// }

	for _, a := range test {
		fmt.Println("\n", a.NAME, a.AKA)
	}
}
*/
