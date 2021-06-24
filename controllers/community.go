package controllers

// -- 与社区相关

import (
	"bulebell/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CommunityHandler 查询分类信息
func CommunityHandler(context *gin.Context) {
	// 1.查询到所有社区
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	// 2.返回响应
	ResponseSuccess(context, data)
}

// CommunityByIDHandler 根据ID查询详细信息
func CommunityDetailHandler(context *gin.Context) {
	// 1.获取ID参数
	idStr := context.Param("id")	// 取到的id 为str类型...虽然最后也可以取到值
	id,err := strconv.ParseInt(idStr,10,64)
	if err != nil {
		ResponseError(context,CodeInvalidParam)
	}

	//2.查找数据
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.CommunityDetailHandler() failed", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	// 3.返回响应
	ResponseSuccess(context, data)
}



