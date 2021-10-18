package entities

import (
	conn "github.com/rpinedafocus/mylib-dbconn"
)

type Authors struct {
	AUTHOR_ID int    `json:"author_id"`
	NAME      string `json:"name"`
}

func Create(newAuthor Authors) bool {

	db, errc := conn.GetConnection()

	if errc {
		return false
	} else {
		sqlStatement := `INSERT INTO university."authors" (name) VALUES ($1)`
		_, err := db.Exec(sqlStatement, newAuthor.NAME)

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
				err := rows.Scan(&author.AUTHOR_ID, &author.NAME)
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

	var test = GetRows()
	// j, err := json.Marshal(test)

	// if err != nil {
	// 	fmt.Printf("Error: %s", err.Error())
	// } else {
	// 	fmt.Println(string(j))
	// }

	for _, a := range test {
		fmt.Println("\n", a.AUTHOR_ID, a.NAME)
	}
}
*/
