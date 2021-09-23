package api

import (
	"net/http"
	"strconv"

	"simple-go-auth/repository/posts"
	"github.com/gin-gonic/gin"
)

func UpdatePost(c *gin.Context) {
	// var user users.User
	// if err := c.ShouldBindJSON(&user); err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, "invalid json")
	// 	return
	// }
	// accessDetails := c.MustGet(auth.ContextAccessDetailsKey).(*auth.AccessDetails)
	// if err := posts.UpdatePost(accessDetails.UserId, user.Email); err != nil {
	// 	c.AbortWithStatusJSON(http.StatusNoContent, err)
	// }

	// return Todo struct
	c.JSON(http.StatusOK, "Updated post info")
}

func GetPost(c *gin.Context) {
	postidStr := c.Param("postid")
	id, _ := strconv.Atoi(postidStr)
	post, err := posts.FindPostByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNoContent, err)
	}
	c.JSON(http.StatusOK, post)
}
