package controllers

import (
	"bulebell/logic"
	"bulebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CreatePostHandler 创建帖子功能
func CreatePostHandler(context *gin.Context) {
	// 1.获取参数以及校验
	p := new(models.Post)
	if err := context.ShouldBindJSON(p); err != nil {
		zap.L().Error("context.ShouldBindJSON() filed", zap.Error(err))
		ResponseError(context, CodeInvalidParam)
		return
	}
	// 2.取到当前发请求的用户ID
	userID, err := getCurrentUser(context)
	if err != nil {
		ResponseError(context, CodeNeedLogin)
		return
	}
	p.AuthorID = userID

	// 3.创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost() is filed", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	// 4.返回响应
	ResponseSuccess(context, nil)
}

// GetPostHandler 获取帖子详情
func GetPostDetailHandler(context *gin.Context) {
	// 1.获取参数
	idStr := context.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(context, CodeInvalidParam)
		return
	}

	// 2.获取数据
	data, err := logic.GetPostDetail(id)
	if err != nil {
		zap.L().Error("logic.GetPostDetail() filed", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	// 3.返回响应
	ResponseSuccess(context, data)
}

// 获取帖子数据
func GetPostListHandler(context *gin.Context) {
	// 获取分页参数
	page, size := getPageInfo(context)
	// 获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList filed", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}

	// 返回响应
	ResponseSuccess(context, data)

}

// GetPostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts2 [get]
func GetPostListHandler2(context *gin.Context) {
	// 获取分页参数
	p := &models.ParamPostList{ // 初始化以及初始参数
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	//context.ShouldBindJSON() 获取json格式的数据
	//context.ShouldBind()	让gin自动动态的做出选择是json还是query
	if err := context.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 is filed", zap.Error(err))
		ResponseError(context, CodeInvalidParam)
		return
	}
	// 获取数据
	data, err := logic.GetPostList2(p)
	if err != nil {
		zap.L().Error("logic.GetPostList filed", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}

	// 返回响应
	ResponseSuccess(context, data)

}

// 根据社区查找帖子
func GetCommunityPostListHandler(context *gin.Context) {
	// 获取分页参数
	p := &models.ParamCommunityPostList{
		ParamPostList: models.ParamPostList{
			Page:  1,
			Size:  10,
			Order: models.OrderTime,
		},
	} // 初始化以及初始参数}
	//context.ShouldBindJSON() 获取json格式的数据
	//context.ShouldBind()	让gin自动动态的做出选择是json还是query
	if err := context.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetCommunityPostListHandler is filed", zap.Error(err))
		ResponseError(context, CodeInvalidParam)
		return
	}
	// 获取数据
	data, err := logic.GetCommunityPostList(p)
	if err != nil {
		zap.L().Error("logic.GetCommunityPostList filed", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}

	// 返回响应
	ResponseSuccess(context, data)
}
