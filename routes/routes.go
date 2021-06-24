package routes

import (
	"bulebell/controllers"
	_ "bulebell/docs"
	"bulebell/logger"
	"bulebell/middlewares"
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func Setup() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)	// 日志是否接收
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true)) //日志的中间件

	// 网页
	r.LoadHTMLFiles("./templates/index.html")
	r.Static("/static", "./static")
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	// 路由组
	v1 := r.Group("/api/v1")

	// 注册路由
	v1.POST("/signup", controllers.SignUpHandler)
	// 登陆功能
	v1.POST("/login", controllers.LoginHandler)

	// 以时间或分数获取帖子列表
	v1.GET("/posts2", controllers.GetPostListHandler2)
	// 根据社区查询帖子列表
	v1.GET("/communityPost", controllers.GetCommunityPostListHandler)
	// 社区分类
	v1.GET("/community", controllers.CommunityHandler)
	// 根据id查询社区详情
	v1.GET("/community/:id", controllers.CommunityDetailHandler)
	// 根据id查询文章详情
	v1.GET("/post/:id", controllers.GetPostDetailHandler)
	// 分页展示帖子数据
	v1.GET("/posts", controllers.GetPostListHandler)

	v1.Use(middlewares.JWTAuthMiddleware()) // 认证jwt中间件

	{
		// 添加文章
		v1.POST("/post", controllers.CreatePostHandler)

		// 投票
		v1.POST("/vote", controllers.PostVoteController)

	}
	// 接口文档
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	// 测试
	//r.GET("/", middlewares.JWTAuthMiddleware(), func(context *gin.Context) {
	//	//如果是登陆的用户，返回ok，（判断请求头中是否有有效的token）
	//	context.String(http.StatusOK, "ok")
	//})

	return r
}
