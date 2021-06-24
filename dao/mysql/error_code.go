package mysql

import "errors"

var (
	ErrUserExist       = errors.New("用户已存在")
	ErrUserNoExist     = errors.New("用户不存在")
	ErrInvalidPassword = errors.New("账号或密码错误")
	ErrInvalidID       = errors.New("无效ID")
)
