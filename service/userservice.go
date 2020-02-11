package service

import (
	"fmt"

	"myapi/handler"
	"myapi/model"
	errno "myapi/pkg/err"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context)  {
	var r model.User
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	u := model.User{
		UserName: r.UserName,
		Password: r.Password,
	}
	// Validate the data.
	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}
	// Insert the user to the database.
	if _,err := u.Create(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	// Show the user information.
	handler.SendResponse(c, nil, u)
}

// SelectUser 查询用户
func SelectUser(c *gin.Context)  {
	name := c.Query("user_name")
	if name == ""{
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}
	var  user model.User
	if err := user.SelectUserByName(name);nil != err {
		fmt.Println(err)
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	// Validate the data.
	if err := user.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	handler.SendResponse(c, nil, user)
}

// 首页
func Index(c *gin.Context)  {
	handler.SendResponse(c, nil, "hello world")
}