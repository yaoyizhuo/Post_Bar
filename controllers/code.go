package controllers

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNoExist
	CodeInvalidPassword
	CodeServerBusy

	//CodeNeedAuth
	CodeInvalidToken
	CodeNeedLogin
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeUserNoExist:     "用户不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",
	//CodeNeedAuth:        "请求头中auth为空",
	CodeInvalidToken:    "无效的Token",
	CodeNeedLogin:       "需要登陆",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
