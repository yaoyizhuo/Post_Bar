package controllers

import (
	"bulebell/logic"
	"bulebell/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 投票的处理函数
func PostVoteController(context *gin.Context) {
	// 1.参数校验
	p := new(models.ParamVoteData)
	err := context.ShouldBindJSON(p)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors) // 类型断言
		if !ok {
			ResponseError(context, CodeInvalidParam)
			return
		}
		errData := errs.Translate(trans)
		ResponseErrorWithMsg(context, CodeInvalidParam, errData)
		return
	}

	// 2.获取用户信息
	userid, err := getCurrentUser(context)
	if err != nil {
		ResponseError(context, CodeNeedLogin)
		return
	}

	// 3.函数处理
	if err := logic.VoteForPost(userid, p); err != nil {
		zap.L().Error("logic.VoteForPost  filed", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	// 4.返回响应
	ResponseSuccess(context, nil)
}
