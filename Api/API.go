package API

import (
	"database/sql"
	"db_p/profile"
	"db_p/signUP_IN"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func setupDB() *sql.DB {
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

	return db
}

type user struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func getID(c *gin.Context) int {
	Id := c.Param("id")
	userId, err := strconv.Atoi(Id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())

		c.AbortWithError(http.StatusInternalServerError, err)
	}
	return userId
}
func CreatUser(c *gin.Context) {
	Email := c.Query("email")
	Password := c.Query("password")
	userId, err := signUP_IN.SignUp(Email, Password, setupDB())
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())

		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": userId, "message": "you signed up successfully!"})
}
func SignInUser(c *gin.Context) {
	Email := c.Query("email")
	Password := c.Query("password")
	userId, err := signUP_IN.SignIn(Email, Password, setupDB())
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())

		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": userId, "message": "you signed up successfully!"})
}

func News(c *gin.Context) {
	Id := getID(c)
	NewProducts := profile.ShowNews(setupDB(), Id)
	c.JSON(http.StatusOK, gin.H{"data": NewProducts, "message": "you signed up successfully!"})
}

func Logs(c *gin.Context) {

}
func ModifyMobile(c *gin.Context) {
	userId := getID(c)
	temp := c.Query("mobile")
	profile.Modify_mobile(userId, temp, setupDB())
}
func ModifyEmail(c *gin.Context) {
	Id := getID(c)
	newEmail := c.Query("email")
	profile.Modify_email(Id, newEmail, setupDB())
}
func ModifyPassword(c *gin.Context) {
	Id := getID(c)
	newPassword := c.Query("password")
	profile.Modify_password(Id, newPassword, setupDB())
}

func ModifyAdd2(c *gin.Context) {
	Id := getID(c)
	add2 := c.Query("address2")
	profile.Modify_address2(Id, add2, setupDB())
}
func ModifyAdd1(c *gin.Context) {
	Id := getID(c)
	add1 := c.Query("address1")
	profile.Modify_address1(Id, add1, setupDB())

}
func ModifyLastName(c *gin.Context) {
	Id := getID(c)
	lName := c.Query("lastName")
	profile.Modify_lastname(Id, lName, setupDB())

}
func ModifyFirstName(c *gin.Context) {
	userId := getID(c)
	temp := c.Query("name")
	profile.Modify_firstname(userId, temp, setupDB())

}
func ProductsByCategory(c *gin.Context) {

}
func AllCategories(c *gin.Context) {

}

func AllProducts(c *gin.Context) {

}
func GetAllProductsByCategory(c *gin.Context) {

}
