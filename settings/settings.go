package settings

import (
	"fmt"
	"github.com/spf13/viper"
)

// 管理配置
func Init() (err error) {
	//viper.SetConfigName("config") // 配置文件名称(无扩展名)
	//viper.SetConfigType("yaml")   // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath("./conf")      // 还可以在工作目录中查找配置

	err = viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {            // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
		return
	}
	return
}
