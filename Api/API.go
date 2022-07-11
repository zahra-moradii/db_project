package API

import (
	"db_p/structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Users []structs.User

func CreateUser(c *gin.Context) {
	var newUser structs.User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	Users = append(Users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}
