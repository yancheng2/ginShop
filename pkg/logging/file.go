package logging

import (
	"fmt"
	"ginShop/pkg/file"
	"ginShop/pkg/setting"
	"os"
	"time"
)
// 获取日志文件路径
func getLogFilePath()string{
	return fmt.Sprintf("%s%s",setting.AppSetting.RuntimeRootPath,setting.AppSetting.LogSavePath)
}

// 获取日志文件的名称
func getLogFileName()string{
	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt,
	)
}

// 打开日志文件
func openLogFile(filename,filepath string)(*os.File,error){
	dir, err := os.Getwd()
	if err!=nil{
		return nil,fmt.Errorf("os.Getwd err:%s",err)
	}
	// 拼接文件路径
	path := fmt.Sprintln(dir,"/",filepath)
	// 检查文件权限
	permission := file.CheckPermission(path)
	if permission == true{
		return nil,fmt.Errorf("file.CheckPermission denied path: %s",path)
	}
	// 如果不存在则新建文件夹
	err = file.IsNotExistMkdir(path)
	if err!=nil{
		return nil,fmt.Errorf("file.IsNotExistMkdir path :%s,err:%v",path,err)
	}
	//调用文件
	file, err := file.Open(path+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err!= nil{
		return nil,fmt.Errorf("fail to openFile :%v",err)
	}
	return file,nil
}