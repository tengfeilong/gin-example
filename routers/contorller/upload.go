package contorller

import (
	"gin-example/mylog"
	"gin-example/utils/msg"
	"gin-example/utils/upload"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func UploadImage(c *gin.Context) {
	code := msg.SUCCESS
	data := make(map[string]string)

	file, image, err := c.Request.FormFile("image")
	if err != nil {
		code = msg.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  msg.GetMsg(code),
			"data": data,
		})
	}

	if image == nil {
		code = msg.INVALID_PARAMS
	} else {
		imageName := upload.GetImageName(image.Filename)
		fullPath := upload.GetImageFullPath()
		savePath := upload.GetImagePath()

		src := fullPath + imageName
		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
			code = msg.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
		} else {
			err := upload.CheckImage(fullPath)
			if err != nil {
				mylog.Logger.Info("upload err", zap.String("err", err.Error()))
				code = msg.ERROR_UPLOAD_CHECK_IMAGE_FAIL
			} else if err := c.SaveUploadedFile(image, src); err != nil {
				mylog.Logger.Info("upload err", zap.String("err", err.Error()))
				code = msg.ERROR_UPLOAD_SAVE_IMAGE_FAIL
			} else {
				data["image_url"] = upload.GetImageFullUrl(imageName)
				data["image_save_url"] = savePath + imageName
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg.GetMsg(code),
		"data": data,
	})
}
