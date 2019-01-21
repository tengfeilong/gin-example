package v1

import (
	"gin-example/config"
	"gin-example/models"
	"gin-example/mylog"
	"gin-example/utils/msg"
	"gin-example/utils/util"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// @Summary 获取多个文章标签
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {string} json "{"code":200,"data":{"lists":[{"id":3,"created_on":1516849721,"modified_on":0,"name":"3333","created_by":"4555","modified_by":"","state":0}],"total":29},"msg":"ok"}"
// @Router /controller/tag/getTagList [get]
func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := msg.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), config.AppSetting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg.GetMsg(code),
		"data": data,
	})
}

//@Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /controller/tag/addTag [post]
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := msg.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = msg.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = msg.ERROR_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			mylog.Logger.Info("tag err", zap.String("key", err.Key), zap.String("err", err.Message))
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg.GetMsg(code),
		"data": make(map[string]string),
	})
}

// @Summary 修改文章标签
// @Produce  json
// @Param id param int true "ID"
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param modified_by query string true "ModifiedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /controller/tag/editTag/{id} [post]
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := msg.INVALID_PARAMS
	if !valid.HasErrors() {
		code = msg.SUCCESS
		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}

			models.EditTag(id, data)
		} else {
			code = msg.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			mylog.Logger.Info("tag err", zap.String("key", err.Key), zap.String("err", err.Message))
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg.GetMsg(code),
		"data": make(map[string]string),
	})
}

// @Summary 删除文章标签
// @Produce  json
// @Param id param int true "ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /controller/tag/deleteTag/{id} [delete]
func DeleteTag(c *gin.Context) {

	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := msg.INVALID_PARAMS
	if !valid.HasErrors() {
		code = msg.SUCCESS
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		} else {
			code = msg.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			mylog.Logger.Info("tag err", zap.String("key", err.Key), zap.String("err", err.Message))
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg.GetMsg(code),
		"data": make(map[string]string),
	})
}
