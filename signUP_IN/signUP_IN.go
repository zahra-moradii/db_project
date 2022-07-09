package signUP_IN

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

func userExist(Email string) (int, bool) {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "ayda",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "mySql",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	var id int

	result, err := db.Query("SELECT user_id FROM user_info WHERE email=?", Email)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
		return -1, false
	}
	for result.Next() {
		err = result.Scan(&id)
	}
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
		return id, false
	}
	return id, true
}
func signUp() error {
	var Email string
	//var Name string

	fmt.Scanln(&Email)
	//todo
	//	fmt.Scanln(&password)

	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "ayda",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "mySql",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	_, err = db.Query("SELECT user_id FROM user_info WHERE email=?", Email)
	if err == nil {
		fmt.Printf("User %s already exists! please signin \n", Email)
		return errors.New("user already exists")
	}

	//todo //add username and get email and ...
	//_, err := db.Query("INSERT INTO `admin_info` (`admin_id`, `admin_name`, `admin_email`, `admin_password`)" +
	//	" VALUES   (1, 'admin', 'admin@gmail.com', '25f9e794323b453885f5181f1b624d0b');\n")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return nil

}
func signIn() (int, error) {
	var email string
	//todo also get password and then check in database
	fmt.Scanln(&email)
	var id int
	//var exist bool
	//id, exist = userExist(email)
	//if exist {
	//todo // go to next page

	//return id, nil
	//	}

	fmt.Printf("User %s does not exist! please signUp \n", email)
	return id, errors.New("user does not exist")

}
