package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-web-collection/config"
)

//定义一个创建文件目录的方法
func FullPath() string {
	folderName := time.Now().Format("2006/01/02")
	folderPath := filepath.Join(config.UploadBasePath(), folderName)
	os.MkdirAll(folderPath, os.ModePerm)
	return folderPath
}

func FullFileName(suffix string) (fullFileName string) {
	fileName := fmt.Sprintf("%d", time.Now().UnixNano())
	fullFileName = filepath.Join(FullPath(), fileName+"."+suffix)
	return
}
