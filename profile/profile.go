package profile

import (
	"database/sql"
	"errors"
	//"fmt"
	//"github.com/go-sql-driver/mysql"
	//"log"
)

func modify_email(id int, Email string, db *sql.DB) {
	sqlStatement := `
		UPDATE user_info
		SET email = $2
		WHERE user_id = $1;`

	_, err := db.Exec(sqlStatement, id, Email)
	if err != nil {
		errors.New("could not change email")
	}
}
func modify_password(id int, Password string, db *sql.DB) {
	sqlStatement := `
		UPDATE user_info
		SET password = $2
		WHERE user_id = $1;`

	_, err := db.Exec(sqlStatement, id, Password)
	if err != nil {
		errors.New("could not change password")
	}
}
func modify_firstname(id int, fname string, db *sql.DB) {
	sqlStatement := `
		UPDATE user_info
		SET first_name = $2
		WHERE user_id = $1;`

	_, err := db.Exec(sqlStatement, id, fname)
	if err != nil {
		errors.New("could not change first_name")
	}
}
func modify_lastname(id int, lname string, db *sql.DB) {
	sqlStatement := `
		UPDATE user_info
		SET last_name = $2
		WHERE user_id = $1;`

	_, err := db.Exec(sqlStatement, id, lname)
	if err != nil {
		errors.New("could not change last_name")
	}
}
func modify_mobile(id int, phnoeNum string, db *sql.DB) {
	sqlStatement := `
		UPDATE user_info
		SET mobile = $2
		WHERE user_id = $1;`

	_, err := db.Exec(sqlStatement, id, phnoeNum)
	if err != nil {
		errors.New("could not change email")
	}
}
func modify_address1(id int, add1 string, db *sql.DB) {
	sqlStatement := `
		UPDATE user_info
		SET address1 = $2
		WHERE user_id = $1;`

	_, err := db.Exec(sqlStatement, id, add1)
	if err != nil {
		errors.New("could not change address1")
	}
}
func modify_address2(id int, add2 string, db *sql.DB) {
	sqlStatement := `
		UPDATE user_info
		SET address2 = $2
		WHERE user_id = $1;`

	_, err := db.Exec(sqlStatement, id, add2)
	if err != nil {
		errors.New("could not change address2")
	}
}
