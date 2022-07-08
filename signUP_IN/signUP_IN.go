package signUP_IN

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func getDataBase() *sql.DB {
	db, err := sql.Open("mysql", "db_name:password@tcp(db:port)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()
	return db
}
func userExist(username string) bool {
	db := getDataBase()
	//todo
	//todo does this username(or email) exist?
	result, err := db.Query("SELECT * FROM")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app

	}
	//todo //check is it right?
	if result != nil {
		return true
	}
	return false
}
func signUp() error {
	var username string
	fmt.Scanln(&username)

	if userExist(username) {
		fmt.Printf("User %s already exists! please signin \n", username)
		return errors.New("user already exists")
	}

	db := getDataBase()

	//todo //add username and get email and ...
	_, err := db.Query("SELECT * FROM")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return nil
}
func signIn() error {
	var username string
	//todo alseo get password and then check in database
	fmt.Scanln(&username)
	if userExist(username) {
		//todo // go to next page

		return nil
	}

	fmt.Printf("User %s does not exist! please signUp \n", username)
	return errors.New("user does not exist")

}
