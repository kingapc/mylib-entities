package manage

import (
	"time"

	conn "github.com/rpinedafocus/mylib-dbconn"
)

type Status struct {
	BOOK_ID            int    `json:"book_id"`
	RESERVED_DATE      string `json:"reserved_date"`
	RENT_DATE          string `json:"rent_date"`
	RETURN_DATE        string `json:"return_date"`
	RENTED_RESERVED_BY string `json:"rented_reserved_by"`
	IS_RETURNED        bool   `json:"is_returned"`
}

func GetReserve(nr Status) bool {

	db, errc := conn.GetConnection()

	if errc {
		return false
	} else {
		dt := time.Now().Local()
		sqlStatement := `INSERT INTO university.status (book_id,reserved_date,rented_reserved_by,is_returned) VALUES ($1,$2,$3,$4)`
		_, err := db.Exec(sqlStatement, nr.BOOK_ID, dt.Format("2006-01-02"), nr.RENTED_RESERVED_BY, nr.IS_RETURNED)

		if err != nil {
			return false
		} else {
			return true
		}
	}
}

func GetRent(nr Status) bool {

	db, errc := conn.GetConnection()

	if errc {
		return false
	} else {
		dt := time.Now().Local()
		sqlStatement := `INSERT INTO university.status (book_id,rent_date,rented_reserved_by,is_returned) VALUES ($1,$2,$3,$4)`
		_, err := db.Exec(sqlStatement, nr.BOOK_ID, dt.Format("2006-01-02"), nr.RENTED_RESERVED_BY, nr.IS_RETURNED)

		if err != nil {
			return false
		} else {
			return true
		}
	}
}

func ReturBook(id int) bool {

	db, errc := conn.GetConnection()

	if errc {
		return false
	} else {
		dt := time.Now().Local()
		sqlStatement := `UPDATE university.status SET return_date = $1, is_returned = true WHERE status_id = $2`
		_, err := db.Exec(sqlStatement, dt.Format("2006-01-02"), id)

		if err != nil {
			return false
		} else {
			return true
		}
	}
}

/*
func main() {

	var a = Status{
		BOOK_ID:            1,
		RENTED_RESERVED_BY: "root",
		IS_RETURNED:        false,
	}

	var test = GetReserve(a)
	fmt.Print(test)
}
*/
