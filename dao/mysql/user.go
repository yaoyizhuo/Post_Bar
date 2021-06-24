package mysql

import (
	"bulebell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
)

const secret = "1106927262@qq.com"

var (
	ErrorUserExist   = errors.New("用户已存在")
	ErrorUserNoExist = errors.New("用户不存在")
	ErrorPassword    = errors.New("密码错误")
)

// CheckUserExist 检查用户是否存在
func CheckUserExist(username string) (bool, error) {
	sqlstr := `select count(user_id) from user where username=?`

	var count int
	err := db.Get(&count, sqlstr, username) // 返回结果，执行语句，参数
	if err != nil {
		return false, ErrorUserExist
	}

	return count > 0, nil
}

// InsertUser 对数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	user.PassWord = encryptPassword(user.PassWord)

	//执行Sql语句入库
	sqlstr := `insert into user(user_id,username,password)values(?,?,?)`
	_, err = db.Exec(sqlstr, user.UserID, user.UserName, user.PassWord)

	return err
}

// GetUser 从数据库中取出用户
func GetUser(user *models.User) (err error) {
	// 执行sql语句
	osPassword := user.PassWord // 原来的密码
	sqlstr := `select user_id,username,password from user where username=?`
	err = db.Get(user, sqlstr, user.UserName)
	if err == sql.ErrNoRows {
		return ErrorUserNoExist
	}
	if err != nil {
		return ErrorUserNoExist
	}

	if encryptPassword(osPassword) != user.PassWord {
		return ErrorPassword
	}

	return
}

// encryptPassword 加密
func encryptPassword(str string) string {
	h := md5.New()
	h.Write([]byte(secret))      // 加盐
	newStr := h.Sum([]byte(str)) // 加密后形成切片

	return hex.EncodeToString(newStr) // 将newStr转化为一个16进制的字符串
}

// GetUserByID 根据作者ID查询作者名
func GetUserByID(uid int64) (user *models.User, err error) {
	user = new(models.User)

	sqlStr := "select user_id,username from user where user_id=?"
	err = db.Get(user, sqlStr, uid)

	return
}
