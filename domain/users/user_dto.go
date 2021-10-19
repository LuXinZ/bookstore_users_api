package users

import (
	"github.com/LuXinZ/bookstore_users_api/utils/errors"
	"strings"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func Validate(user *User) *errors.RestErr  {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email =="" {
		return errors.NewBadRequestError("无效的邮箱")
	}
	return nil
}
func (user *User) Validate() *errors.RestErr  {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email =="" {
		return errors.NewBadRequestError("无效的邮箱")
	}
	return nil
}