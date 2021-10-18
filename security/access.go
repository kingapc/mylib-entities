package security

import (
	conn "github.com/rpinedafocus/mylib-dbconn"
)

type Users struct {
	USER_ID    int    `json:"user_id"`
	USER_NAME  string `json:"user_name"`
	FIRST_NAME string `json:"first_name"`
	LAST_NAME  string `json:"last_name"`
	EMAIL      string `json:"email"`
	ROLE_ID    string `json:"role_id"`
	PASSWORD   string `json:"password"`
}

type InfoLogin struct {
	FULL_NAME string `json:"full_name"`
	USER_NAME string `json:"user_name"`
	ROLE_ID   int    `json:"role_id"`
}

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

func Login(user string, password string) (InfoLogin, bool) {

	db, errc := conn.GetConnection()

	var infoLogin InfoLogin

	if errc {
		return infoLogin, true
	}

	stmt, err := db.Prepare(`SELECT CONCAT(trim(first_name),' ',trim(last_name)) as full_name, user_name as user, role_id level FROM university.users WHERE trim(user_name) = $1 AND trim(password) = $2`)
	if err != nil {
		return infoLogin, true
	}
	err = stmt.QueryRow(user, password).Scan(&infoLogin.FULL_NAME, &infoLogin.USER_NAME, &infoLogin.ROLE_ID)

	return infoLogin, false
}

func CreateUser(nu Users) (Users, bool) {

	db, errc := conn.GetConnection()

	if errc {
		return nu, false
	} else {
		sqlStatement := `INSERT INTO university.users (user_name, first_name, last_name, email, role_id, password) VALUES ($1,$2,$3,$4,$5,$6)`
		_, err := db.Exec(sqlStatement, nu.USER_NAME, nu.FIRST_NAME, nu.LAST_NAME, nu.EMAIL, nu.ROLE_ID, nu.PASSWORD)

		if err != nil {
			return nu, false
		} else {
			return nu, true
		}
	}
}

/*
func main() {

	test, err := Login("root", "root")

	if err {
		fmt.Println("you don't have access")
	} else {
		fmt.Print(test)
	}
}
*/
