package file

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

//获取文件大小
func GetSize(f multipart.File)(int,error){
	context, err := ioutil.ReadAll(f)
	return len(context),err
}
//获取文件后缀
func GetExt(filename string) string{
	return path.Ext(filename)
}
//检查文件是否存在
/*
   如果返回的错误为nil,说明文件或文件夹存在
   如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
   如果返回的错误为其它类型,则不确定是否在存在
*/
func CheckExist(path string)bool{
	_, err := os.Stat(path)
	if err == nil{
		return true
	}
	if os.IsNotExist(err){
		return false
	}
	return false
}

//检查文件权限  true-无权限  false-有权限
func CheckPermission(path string)bool{
	_, err := os.Stat(path)
	return os.IsPermission(err)
}

//新建文件夹
func Mkdir(path string)error{
	err := os.MkdirAll(path, os.ModePerm)
	return err
}

//如果不存在则新建文件夹
func IsNotExistMkdir(path string)error{
	if exist := CheckExist(path);exist==false{
		if err := Mkdir(path);err!=nil{
			return err
		}
	}
	return nil
}
/*
   调用文件，支持传入文件名称、指定的模式调用文件、文件权限，返回的文件的方法可以用于I/O。如果出现错误，则为*PathError
   const (
       // Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
       O_RDONLY int = syscall.O_RDONLY // 以只读模式打开文件
       O_WRONLY int = syscall.O_WRONLY // 以只写模式打开文件
       O_RDWR   int = syscall.O_RDWR   // 以读写模式打开文件
       // The remaining values may be or'ed in to control behavior.
       O_APPEND int = syscall.O_APPEND // 在写入时将数据追加到文件中
       O_CREATE int = syscall.O_CREAT  // 如果不存在，则创建一个新文件
       O_EXCL   int = syscall.O_EXCL   // 使用O_CREATE时，文件必须不存在
       O_SYNC   int = syscall.O_SYNC   // 同步IO
       O_TRUNC  int = syscall.O_TRUNC  // 如果可以，打开时
   )
*/
func Open(name string,flag int,perm os.FileMode)(*os.File,error){
	file, err := os.OpenFile(name, flag, perm)
	if err!=nil{
		return nil,err
	}
	return file,err
}