package cgo

import (
	"log"
	"os"
	"io"
	"cgo/utils"
	"net/http"
	"cgo/constant"
)

type Controller struct {
	Data interface{}
}

type FileInfoTO struct {
	//图片id -- 暂时没有用
	ID int64
	//缩略图路径 -- 暂时没有用
	CompressPath string
	//原图路径 ,保存数据库的路径
	Path string
	//原始的文件名
	OriginalFileName string
	//存储文件名 如：uuidutil
	FileName string
	//文件大小
	FileSize int64
}

//解析Form-data中的文件，不管上传的文件的字段名(fieldname)是什么，都会解析
func (p *Controller) SaveFiles(r *http.Request,relativePath string) []*FileInfoTO {
	r.ParseMultipartForm(32 << 20)
	m := r.MultipartForm
	if m == nil {
		log.Println("not multipartfrom !")
		return nil
	}
	fileInfos := make([]*FileInfoTO,0)

	filePath := constant.BASE_IMAGE_ADDRESS + relativePath
	utils.MakeDir(filePath)

	//files := m.File["files"] //根据上传文件时指定的字段名(fieldname)获取FileHeaders
	for _,fileHeaders := range m.File { //遍历所有的所有的字段名(fieldname)获取FileHeaders
		for _,fileHeader := range fileHeaders{
			file,err := fileHeader.Open()
			if err != nil {
				log.Println(err)
				return fileInfos
			}
			defer file.Close()
			name,err := utils.RandomUUID()
			if err != nil {
				log.Println(err)
				return fileInfos
			}
			fileType := utils.Ext(fileHeader.Filename,".jpg")
			newName := name.String() + fileType
			dst,err := os.Create(filePath + newName)
			if err != nil {
				log.Println(err)
				return fileInfos
			}
			fileSize,err := io.Copy(dst,file)
			if err != nil {
				log.Println(err)
				return fileInfos
			}
			fileInfos = append(fileInfos, &FileInfoTO{Path:relativePath + newName,OriginalFileName:fileHeader.Filename,FileName:newName,FileSize:fileSize})
		}
	}
	return fileInfos
}
