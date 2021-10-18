package security

import (
	conn "github.com/rpinedafocus/mylib-dbconn"
)

func GetAccess(myLevel int, endpoint string) bool {

	db, errc := conn.GetConnection()

	if errc {
		return false
	} else {
		var level int
		stmt, err := db.Prepare(`SELECT level FROM university.access WHERE endpoint = $1`)
		err = stmt.QueryRow(endpoint).Scan(&level)

		if err != nil {
			return false
		} else {
			if myLevel == level {
				return true
			} else if level == 3 {
				return true
			} else {
				return false
			}
		}
	}
}

/*
func main() {

	var test = GetAccess(3, "/books")

	if test {
		fmt.Println("you have access")
	} else {
		fmt.Println("you don't have access")
	}
}
*/
