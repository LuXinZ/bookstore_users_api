package users

import (
	"github.com/LuXinZ/bookstore_users_api/domain/users"
	"github.com/LuXinZ/bookstore_users_api/services"
	"github.com/LuXinZ/bookstore_users_api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	userId, userErr := strconv.ParseInt(c.Param("user_id"),10,64)
	if userErr != nil {
		err := errors.NewBadRequestError("用户ID 应该是一个数字")
		c.JSON(err.Status,err)
		return
	}
	user, getError := services.GetUser(userId)
	if getError != nil {
		c.JSON(getError.Status,getError)
		return
	}
	c.JSON(http.StatusOK,user)

}

func SearchUser(c *gin.Context)  {
	c.String(http.StatusNotImplemented,"implement me")
}