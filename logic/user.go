package logic

import (
	"bulebell/dao/mysql"
	"bulebell/models"
	"bulebell/pkg/jwt"
	"bulebell/pkg/snowflake"
	"errors"
)

var exist bool // 判断用户是否存在

// 注册
func SignUp(p *models.ParamsSignUp) (err error) {
	//1.判断用户是否存在

	exist, err = mysql.CheckUserExist(p.Username)
	if err != nil {
		// 查询出错
		return err
	}
	if exist {
		return errors.New("用户已存在")
	}

	//2.生成UID
	userID := snowflake.GetID()
	//构造一个User实例
	user := &models.User{
		UserID:   userID,
		UserName: p.Username,
		PassWord: p.Password,
	}

	//3.保存进数据库
	err = mysql.InsertUser(user)
	if err != nil {
		return errors.New("保存数据库失败！")
	}

	return
	//redis...
}

// 登陆
func LoginUp(l *models.ParamsLoginUser) (user *models.User, err error) {
	//1.根据用户名取出用户
	user = &models.User{
		UserName: l.Username,
		PassWord: l.Password,
	}
	err = mysql.GetUser(user)
	if err != nil {
		return nil, errors.New("提取用户失败")
	}

	// 4.生成JWT
	token,err := jwt.GenToken(user.UserID, user.UserName)
	if err != nil {
		return nil, errors.New("生成JWT失败")
	}
	user.Token = token
	return
}
