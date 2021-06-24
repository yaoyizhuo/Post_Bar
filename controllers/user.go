package controllers

// -- 与用户注册登陆相关

import (
	"bulebell/dao/mysql"
	"bulebell/logic"
	"bulebell/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// SignUpHandler 登陆账号
// @Summary 登陆账号
// @Description 可按账号和密码进行账号登陆
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "账号 密码"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts2 [get]
func SignUpHandler(context *gin.Context) {
	//1.获取参数并校验
	p := new(models.ParamsSignUp)
	err := context.ShouldBindJSON(p)
	if err != nil {
		// 请求参数有误
		zap.L().Error("SignUp获取参数失败", zap.Error(err))
		// 判断参数是否是翻译器所需要的类型，如果不是，没有必要翻译
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(context, CodeInvalidParam)
			return
		}
		// 翻译出错
		ResponseErrorWithMsg(context, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 手动对参数进行校验，不要相信前端的js，也防止了禁用js的攻击(过于麻烦，会有大量重复的校验)
	//if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.RePassword != p.Password {
	//	// 请求参数有误
	//	zap.L().Error("SignUp获取参数失败")
	//	context.JSON(http.StatusOK, gin.H{
	//		"message": "请求参数有误",
	//	})
	//	return
	//}

	//2.业务处理
	if err := logic.SignUp(p); err != nil {

		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(context, CodeUserExist)
			return
		}
		ResponseError(context, CodeServerBusy)
		return
	}

	//3.返回响应
	ResponseSuccess(context, nil)
}

func LoginHandler(context *gin.Context) {
	// 1. 获取参数
	l := new(models.ParamsLoginUser)
	err := context.ShouldBindJSON(&l)
	//err := context.ShouldBind(&l)
	if err != nil {
		// 请求参数有误
		zap.L().Error("Login获取参数失败", zap.Error(err))
		// 判断参数是否是翻译器所需要的类型，如果不是，没有必要翻译
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(context, CodeInvalidParam)
			return
		}

		ResponseErrorWithMsg(context, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return

	}
	//fmt.Println(l)

	// 2. 校验参数
	user, err := logic.LoginUp(l)
	if err != nil {
		if errors.Is(err, mysql.ErrorUserNoExist) {
			ResponseError(context, CodeUserNoExist)
			return
		}
		ResponseError(context, CodeInvalidPassword)
		return
	}

	// 3.返回响应
	ResponseSuccess(context, gin.H{
		"user_id":   user.UserID,
		"user_name": user.UserName,
		"token":     user.Token,
	})

}
