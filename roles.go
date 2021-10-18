package entities

import (
	conn "github.com/rpinedafocus/mylib-dbconn"
)

type Roles struct {
	role_name string
}

func CreateRole(newRole Roles) bool {

	db, errc := conn.GetConnection()

	if errc {
		return false
	} else {
		sqlStatement := `INSERT INTO university."roles" (role_name) VALUES ($1)`
		_, err := db.Exec(sqlStatement, newRole.role_name)

		if err != nil {
			return false
		} else {
			return true
		}
	}
}

/*
func main() {
	test := Roles{}
	test.role_name = "Test"

	result := Create(test)

	if result {
		fmt.Printf("\nRegistro creado\n")
	} else {
		fmt.Printf("\nRegistro no creado\n")
	}
}
*/
