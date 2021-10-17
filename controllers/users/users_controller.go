package users

import (
	"github.com/LuXinZ/bookstore_users_api/domain/users"
	"github.com/LuXinZ/bookstore_users_api/services"
	"github.com/LuXinZ/bookstore_users_api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)
func CreateUser(c *gin.Context)  {
	var user users.User
	if err := c.ShouldBindJSON(&user); err!=nil {
		restErr :=errors.NewBadRequestError("错误的JSON body 格式")
		c.JSON(restErr.Status,restErr)
		return
	}
	result, saveError := services.CreateUser(user)
	if saveError != nil {
		c.JSON(saveError.Status,saveError)
		return
	}
	c.JSON(http.StatusCreated,result)
}
func GetUser(c *gin.Context)  {
	c.String(http.StatusNotImplemented,"implement me")
}

func SearchUser(c *gin.Context)  {
	c.String(http.StatusNotImplemented,"implement me")
}