package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

const CtxtUserIDKey = "userID"

var ErrorUserNotLogin = errors.New("用户未登录")

// getCurrentUser 获取当前登陆的用户ID
func getCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxtUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}

	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

// 获取分页的数值
func getPageInfo(c *gin.Context)(int64,int64) {
	// 获取分页参数
	var (
		page int64
		size int64
		err  error
	)
	pageStr := c.Query("page")
	sizeStr := c.Query("size")
	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}

	return page,size
}
