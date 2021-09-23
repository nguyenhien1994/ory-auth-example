package api

import (
	"net/http"
	"strconv"

	"simple-go-auth/repository/users"
	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	// userid := c.Param("userid")

	// return Todo struct
	c.JSON(http.StatusOK, "Updated user info")
}

func GetUser(c *gin.Context) {
	useridStr := c.Param("userid")
	id,_ := strconv.Atoi(useridStr)
	user, err := users.FindUserByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNoContent, err)
	}

	c.JSON(http.StatusOK, user)
}
