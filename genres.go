package entities

import (
	conn "github.com/rpinedafocus/mylib-dbconn"
)

type Genres struct {
	name string
}

func CreateGenre(newGenre Genres) bool {

	db, errc := conn.GetConnection()

	if errc {
		return false
	} else {
		sqlStatement := `INSERT INTO university."genres" (name) VALUES ($1)`
		_, err := db.Exec(sqlStatement, newGenre.name)

		if err != nil {
			return false
		} else {
			return true
		}
	}
}

/*
func main() {
	test := Genres{}
	test.name = "Test Genre"

	result := Create(test)

	if result {
		fmt.Printf("\nRegistro creado\n")
	} else {
		fmt.Printf("\nRegistro no creado\n")
	}
}
*/
