package profile

import (
	"database/sql"
	"errors"
	//"fmt"
	//"github.com/go-sql-driver/mysql"
	//"log"
)

//todo log out

//todo is this an admin or not   profiles are different

func modify_email(id int, Email string, db *sql.DB) {
	_, err := db.Query(`UPDATE user_info SET email = ? WHERE user_id = ?`, Email, id)
	if err != nil {
		errors.New("could not change email")
	}
}
func modify_password(id int, Password string, db *sql.DB) {

	_, err := db.Query(`UPDATE user_info SET password = ? WHERE user_id = ?`, Password, id)
	if err != nil {
		errors.New("could not change password")
	}
}
func modify_firstname(id int, fname string, db *sql.DB) {

	_, err := db.Query(`UPDATE user_info SET first_name = ?	WHERE user_id = ?`, fname, id)
	if err != nil {
		errors.New("could not change first_name")
	}
}
func modify_lastname(id int, lname string, db *sql.DB) {
	_, err := db.Exec(`UPDATE user_info SET last_name = ? WHERE user_id = ?`, lname, id)
	if err != nil {
		errors.New("could not change last_name")
	}
}
func modify_mobile(id int, phnoeNum string, db *sql.DB) {
	_, err := db.Query(`UPDATE user_info SET mobile = ? WHERE user_id = ?;`, phnoeNum, id)
	if err != nil {
		errors.New("could not change email")
	}
}
func Modify_address1(id int, add1 string, db *sql.DB) {

	_, err := db.Query(`UPDATE user_info SET address1 = ? WHERE user_id = ?`, add1, id)
	if err != nil {
		errors.New("could not change address1")
	}
}
func Modify_address2(id int, add2 string, db *sql.DB) {

	_, err := db.Query(`UPDATE user_info SET address2 = ? WHERE user_id = ?`, add2, id)
	if err != nil {
		errors.New("could not change address2")
	}
}
