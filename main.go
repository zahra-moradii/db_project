package main

import (
	"database/sql"
	"db_p/signUP_IN"
	"fmt"

	//	"github.com/aydaZaman/db_project/signUP_IN"
	//"db_p/signUP_IN"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	//	"project/mySql"
	//	"errors"
	//"fmt"
	//	"github.com/gin-gonic/gin"
	"log"
	//	"net/http"
	//signInUp "project/db_project/signUP_IN"
	//	structs "project/db_project/structs"
)

/*
func Users() []*structs.User {
	// Open up our database connection.
	db, err := sql.Open("mysql", "db_name:password@tcp(db:port)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var users []*structs.User
	for results.Next() {
		var u structs.User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&u.ID, &u.Username)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		users = append(users, &u)
	}

	return users
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Users())
}
*/

func getDatabase() *sql.DB {
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
	return db
}

func main() {
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

	id, err := signUP_IN.SignIn(db)
	//err = SignUp(db)
	if err != nil {
		panic(err)
	}
	println("yes")
	fmt.Printf("%d", id)
}
