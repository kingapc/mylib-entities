package entities

import (
	conn "github.com/rpinedafocus/mylib-dbconn"
)

type Authors struct {
	author_id int    `json:"author_id"`
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
		rows, err := db.Query("SELECT author_id, trim(name) FROM university.authors")

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

			// var auth1 = []Authors{
			// 	{author_id: 1, name: "test1"},
			// 	{author_id: 2, name: "test2"},
			// }
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

	var test = GetRows()
	// j, err := json.Marshal(test)

	// if err != nil {
	// 	fmt.Printf("Error: %s", err.Error())
	// } else {
	// 	fmt.Println(string(j))
	// }

	for _, a := range test {
		fmt.Println("\n", a.author_id, a.name)
	}
}
*/
