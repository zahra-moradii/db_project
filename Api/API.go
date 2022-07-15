package API

import (
	"database/sql"
	"db_p/pickbuy"
	"db_p/profile"
	"db_p/signUP_IN"
	"db_p/structs"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
)

func checkErr(c *gin.Context, err error, x gin.H) {
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())

		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, x)
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
	checkErr(c, err, gin.H{"data": userId})
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
	c.JSON(http.StatusOK, gin.H{"data": userId, "email": Email, "password": Password, "message": "you signed up successfully!"})
}
func SignInUser(c *gin.Context) {
	Email := c.Param("email")
	Password := c.Param("password")
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
	c.JSON(http.StatusOK, gin.H{"data": NewProducts, "message": "products that were unavailable for you are available now !"})
}
func ModifyMobile(c *gin.Context) {
	userId := getID(c)
	temp := c.Query("mobile")
	err := profile.Modify_mobile(userId, temp, setupDB())
	checkErr(c, err, gin.H{"data": userId, "message": "you changed your phone number!"})

}
func ModifyEmail(c *gin.Context) {
	Id := getID(c)
	newEmail := c.Query("email")
	err := profile.Modify_email(Id, newEmail, setupDB())
	checkErr(c, err, gin.H{"data": Id, "message": "you changed your Email!"})
}
func ModifyPassword(c *gin.Context) {
	Id := getID(c)
	newPassword := c.Query("password")
	err := profile.Modify_password(Id, newPassword, setupDB())
	checkErr(c, err, gin.H{"data": Id, "message": "you changed your password!"})
}

func ModifyAdd2(c *gin.Context) {
	Id := getID(c)
	add2 := c.Query("address2")
	err := profile.Modify_address2(Id, add2, setupDB())
	checkErr(c, err, gin.H{"data": Id, "message": "you changed your address2!"})
}
func ModifyAdd1(c *gin.Context) {
	Id := getID(c)
	add1 := c.Query("address1")
	err := profile.Modify_address1(Id, add1, setupDB())
	checkErr(c, err, gin.H{"data": Id, "message": "you changed your address1!"})

}
func ModifyLastName(c *gin.Context) {
	Id := getID(c)
	lName := c.Query("lastName")
	err := profile.Modify_lastname(Id, lName, setupDB())
	checkErr(c, err, gin.H{"data": Id, "message": "you changed your last name!"})
}
func ModifyFirstName(c *gin.Context) {
	userId := getID(c)
	temp := c.Query("name")
	err := profile.Modify_firstname(userId, temp, setupDB())
	checkErr(c, err, gin.H{"data": userId, "message": "you changed your first name!"})
}
func ProductsByCategory(c *gin.Context) {
	catId := c.Param("catId")

	id, err := strconv.Atoi(catId)
	checkErr(c, err, gin.H{"data": id})
	db := setupDB()
	products, err := pickbuy.GetProductsByCat(id, db)
	checkErr(c, err, gin.H{"products": products, "category": catId, "message": "all products by category"})
}
func Logs(c *gin.Context) {
	Id := getID(c)
	var logs []structs.Logs
	logs, err := pickbuy.GetLogs(setupDB(), Id)

	checkErr(c, err, gin.H{"logs": logs, "userId": Id, "message": "your log history!"})
}

func AllCategories(c *gin.Context) {
	db := setupDB()
	allCat, err := pickbuy.GetAllCategories(db)
	checkErr(c, err, gin.H{"categories": allCat, "message": "All categories"})

}

func GetAllProductsByCategory(c *gin.Context) {
	db := setupDB()
	allCat, err := pickbuy.GetAllCategories(db)
	checkErr(c, err, gin.H{"categories": allCat, "message": "All categories"})

	var cat structs.Categories
	for cat = range allCat {
		id := cat.Cat_id
		db := setupDB()
		products, err := pickbuy.GetProductsByCat(id, db)
		checkErr(c, err, gin.H{"products": products, "category": id, "message": "all products by all category"})

	}

}
