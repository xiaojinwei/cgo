package utils

import (
	"log"
	"os"
	"path"
)

//创建多级目录
func MkDirAll(path string) bool {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

//检测文件夹或文件是否存在
func Exist(file string) bool {
	if _,err := os.Stat(file);os.IsNotExist(err){
		return false
	}
	return true
}

//获取文件的类型，如：.jpg
//如果获取不到，返回默认类型defaultExt
func Ext(fileName string,defaultExt string) string {
	t := path.Ext(fileName)
	if len(t) == 0 {
		return defaultExt
	}
	return t
}

/// 检验文件夹是否存在，不存在 就创建
func MakeDir(filePath string){
	if !Exist(filePath) {
		MkDirAll(filePath)
	}
}
