package entities

import (
	conn "github.com/rpinedafocus/mylib-dbconn"
)

type Authors struct {
	author_id int    `json:"id"`
	name      string `json:"name"`
}

func Create(newAuthor Authors) bool {

	db, errc := conn.GetConnection()

	if errc {
		return false
	} else {
		sqlStatement := `INSERT INTO university."authors" (name) VALUES ($1)`
		_, err := db.Exec(sqlStatement, newAuthor.name)

		if err != nil {
			return false
		} else {
			return true
		}
	}
}

func GetRows() []Authors {

	var author Authors
	var auth []Authors
	db, errc := conn.GetConnection()

	if errc {
		return auth
	} else {
		rows, err := db.Query("SELECT author_id, name FROM university.authors")

		if err != nil {
			return auth
		} else {
			for rows.Next() {
				err := rows.Scan(&author.author_id, &author.name)
				if err != nil {
					return auth
				}

				auth = append(auth, author)
			}

			return auth
		}
	}
}

/*
func main() {
	// test := Authors{}
	// test.name = "Author Test"

	// result := Create(test)

	// if result {
	// 	fmt.Printf("\nRegistro creado\n")
	// } else {
	// 	fmt.Printf("\nRegistro no creado\n")
	// }

	test := GetRows()

	for _, a := range test {
		fmt.Println("\n", a.author_id, a.name)
	}
}
*/
