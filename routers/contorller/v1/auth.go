package v1

import (
	"gin-example/models"
	"gin-example/utils/msg"
	"gin-example/utils/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type auth struct {
	Username string `form:"username" json:"username" valid:"Required; MaxSize(50)"`
	Password string `form:"password" json:"password" valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	//username := c.Query("username")
	//password := c.Query("password")
	a := auth{}
	c.Bind(&a)
	valid := validation.Validation{}
	//a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(a)

	data := make(map[string]interface{})
	code := msg.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(a.Username, a.Password)
		if isExist {
			token, err := util.GenerateToken(a.Username, a.Password)
			if err != nil {
				code = msg.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token

				code = msg.SUCCESS
			}

		} else {
			code = msg.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg.GetMsg(code),
		"data": data,
	})
}
