package main

import (
	"bulebell/controllers"
	"bulebell/dao/mysql"
	"bulebell/dao/redis"
	"bulebell/logger"
	"bulebell/pkg/snowflake"
	"bulebell/routes"
	"bulebell/settings"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	//加载配置文件，读取config
	if err := settings.Init(); err != nil {
		fmt.Println("init settings failed")
		return
	}

	//初始化日志
	if err := logger.Init(viper.GetString("app.mode")); err != nil {
		fmt.Println("init logger failed")
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("log init success")

	//初始化MySql
	if err := mysql.Init(); err != nil {
		fmt.Println("init mysql failed")
		return
	}
	defer mysql.Close()

	//初始化Redis
	if err := redis.Init(); err != nil {
		fmt.Println("init redis failed")
		return
	}
	defer redis.Close()

	// 初始化雪花算法
	if err := snowflake.Init(viper.GetString("app.start_time"), viper.GetInt64("app.machine_id")); err != nil {
		fmt.Println("init snowflake failed")
		return
	}

	// 初始化翻译器
	err := controllers.InitTrans("zh")
	if err != nil {
		fmt.Println("init Trans failed")
		return
	}


	//注册路由
	r := routes.Setup()
	err = r.Run(fmt.Sprintf(":%d",viper.GetInt("app.port")))
	if err != nil {
		fmt.Println("run server failed")
		return
	}

	//启动服务（优雅关机的方式）
	//srv := &http.Server{
	//	Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
	//	Handler: r,
	//}
	//
	//go func() {
	//	// 开启一个goroutine启动服务
	//	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	//		log.Fatalf("listen: %s\n", err)
	//	}
	//}()
	//
	//// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	//quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	//// kill 默认会发送 syscall.SIGTERM 信号
	//// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	//// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	//// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	//<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	//zap.L().Info("Shutdown Server ...")
	//// 创建一个5秒超时的context
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	//if err := srv.Shutdown(ctx); err != nil {
	//	zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	//}

	zap.L().Info("Server exiting")
}
