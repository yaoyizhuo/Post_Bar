package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Debug("connect mysql failed...", zap.Error(err))
		return
	}

	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"), ) // 最大链接数
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"), ) // 最大空闲链接数
	return
}

func Close() {
	_ = db.Close()
}
