package models

import (
	"fmt"
	"ginShop/pkg/logging"
	"ginShop/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func SetUp(){
	var (
		err error
		databaseType = setting.DatabaseSetting.Type
		user 	     = setting.DatabaseSetting.User
		password 	 = setting.DatabaseSetting.Password
		host 	  	 = setting.DatabaseSetting.Host
		name 		 = setting.DatabaseSetting.Name
	)
	connect := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True",user,password,host,name)
	db, err = gorm.Open(databaseType, connect)
	if err != nil{
		logging.Fatal("数据库连接失败",err)  //数据库连接失败为致命错误
	}

	//设置表名称的前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true)		//设置禁用表名的复数形式
	db.LogMode(true)				//打印sql日志

	db.DB().SetMaxIdleConns(10)     //设置空闲时的最大连接数
	db.DB().SetMaxOpenConns(100)        //设置数据库的最大打开连接数
}
