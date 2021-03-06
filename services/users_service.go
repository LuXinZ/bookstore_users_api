package services

import (
	"github.com/LuXinZ/bookstore_users_api/domain/users"
	"github.com/LuXinZ/bookstore_users_api/utils/errors"
)

func GetUser(userID int64)(*users.User, *errors.RestErr)  {
	result := &users.User{Id: userID}
	if err := result.Get(); err!=nil {
		return nil,err
	}
	return result,nil
}
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err!=nil {
		return nil,err
	}
	if err := user.Save(); err!=nil{
		return nil,err
	}
	return  &user, nil
}