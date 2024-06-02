package global

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"new-go-edas-admin/models/system"
)

var (
	GORM *gorm.DB
	err  error
)

// 初始化数据库
func InitMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.Dbuser"),
		viper.GetString("mysql.DbPwd"),
		viper.GetString("mysql.DbHost"),
		viper.GetInt("mysql.DbPort"),
		viper.GetString("mysql.DbName"),
	)
	GORM, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Println(dsn)
		fmt.Println(viper.GetString("mysql.Dbuser"))
		panic("数据库连接失败")
	}
	if viper.GetInt("mysql.ActiveDebug") == 1 {
		GORM = GORM.Debug()
	}
	//开启连接
	db, _ := GORM.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	TPLogger.Info("数据库初始化成功!!!")
}

// 初始化数据库表
func InitMysqlTables() {
	err = GORM.AutoMigrate(
		system.OperationLog{},
		system.User{},
		system.Menu{},
		system.Role{},
		system.Dept{},
		system.APIPath{},
	)
}
