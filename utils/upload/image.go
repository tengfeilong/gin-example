package upload

import (
	"fmt"
	"gin-example/config"
	"gin-example/utils/util"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

//获取图片完整访问URL
func GetImageFullUrl(name string) string {
	return config.AppSetting.ImagePrefixUrl + "/" + GetImagePath() + name
}

//获取图片md5名称
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

//获取图片路径
func GetImagePath() string {
	return config.AppSetting.ImageSavePath
}

//获取图片完整路径
func GetImageFullPath() string {
	return config.AppSetting.RuntimeRootPath + GetImagePath()
}

//检查图片后缀
func CheckImageExt(fileName string) bool {
	ext := GetExt(fileName)
	for _, allowExt := range config.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

//检查图片大小
func CheckImageSize(f multipart.File) bool {
	size, err := GetSize(f)
	if err != nil {
		log.Println(err)
		return false
	}

	return size <= config.AppSetting.ImageMaxSize
}

//检查图片
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
