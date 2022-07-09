package signUP_IN

import (
	"database/sql"
	"errors"
	"fmt"
)

func userExist(Email string, db *sql.DB) (int, bool) {
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
func signUp(db *sql.DB) error {
	var Email string
	var Password string

	fmt.Println("email:")
	fmt.Scanln(&Email)
	fmt.Println("password")
	fmt.Scanln(&Password)

	_, err := db.Query("SELECT user_id FROM user_info WHERE email=?", Email)
	if err == nil {
		fmt.Printf("User %s already exists! please signin \n", Email)
		return errors.New("user already exists")
	}

	//add new user
	sqlStatement := `
	 INSERT INTO user_info
	 SET email = $1, password = $2 ,first_name =$3
	,last_name=$4,mobile=$5,address1=$6,address2=$7	;`

	_, err = db.Exec(sqlStatement, Email, Password, " ", " ", " ", " ", " ")

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return nil

}
func signIn(db *sql.DB) (int, error) {
	var Email string
	var Password string
	// get email, password and then check in database

	fmt.Println("email:")
	fmt.Scanln(&Email)
	fmt.Println("password")
	fmt.Scanln(&Password)

	var id int
	var exist bool
	id, exist = userExist(Email, db)
	if exist {
		//check password
		_, err := db.Query("SELECT user_id FROM user_info WHERE password=?", Password)
		if err == nil {
			fmt.Printf("password is incorrect!\ntry again.")
		}
		return id, nil
	}

	fmt.Printf("User %s does not exist! please signUp \n", Email)
	return id, errors.New("please signup")

}
