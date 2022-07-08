package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	structs "project/db_project/structs"
	//"github.com/gin-gonic/gin"
)

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

func main() {
	fmt.Println("User already exists! please signin \n")
	//router := gin.Default()
	//router.GET("/users", getUsers)

	//router.Run("localhost:8080")

}
