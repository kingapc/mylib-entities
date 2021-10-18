package manage

import (
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
		sqlStatement := `INSERT INTO university.status (book_id,reserved_date,rent_date,return_date,rented_reserved_by,is_returned) VALUES ($1,$2,$3,$4,$5,$6)`
		_, err := db.Exec(sqlStatement, nr.BOOK_ID, nr.RESERVED_DATE, nr.RENT_DATE, nr.RETURN_DATE, nr.RENTED_RESERVED_BY, nr.IS_RETURNED)

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
		sqlStatement := `INSERT INTO university.status (book_id,reserved_date,rent_date,return_date,rented_reserved_by,is_returned) VALUES ($1,$2,$3,$4,$5,$6)`
		_, err := db.Exec(sqlStatement, nr.BOOK_ID, nr.RESERVED_DATE, nr.RENT_DATE, nr.RETURN_DATE, nr.RENTED_RESERVED_BY, nr.IS_RETURNED)

		if err != nil {
			return false
		} else {
			return true
		}
	}
}
