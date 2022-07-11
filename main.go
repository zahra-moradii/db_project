package main

import (
	"database/sql"
	"db_p/pickbuy"
	"db_p/signUP_IN"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"log"
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
*/

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

	//router := gin.Default()
	//router.POST("/users", API.CreateUser)

	id, _ := signUP_IN.SignIn(db)
	pickbuy.Order(db, id)
	pickbuy.Buy(db, id)
}
