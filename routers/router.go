package routers

import (
	"gin-example/config"
	_ "gin-example/docs"
	"gin-example/routers/contorller"
	"gin-example/routers/contorller/v1"
	"gin-example/utils/upload"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func InitRouter() *gin.Engine {
	rou := gin.New()
	rou.Use(gin.Logger(), gin.Recovery())
	gin.SetMode(config.ServerSetting.RunMode)

	rou.POST("/auth", v1.GetAuth)
	rou.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	rou.POST("/upload", contorller.UploadImage)
	//文件访问
	rou.StaticFS("/images", http.Dir(upload.GetImageFullPath()))
	//操作标签
	contorllerTag := rou.Group("/controller/tag")
	//contorllerTag.Use(middleware.JWT())
	{
		//获取标签列表
		contorllerTag.GET("/getTagList", v1.GetTags)
		//新建标签
		contorllerTag.POST("/addTag", v1.AddTag)
		//更新指定标签
		contorllerTag.POST("/editTag/:id", v1.EditTag)
		//删除指定标签
		contorllerTag.DELETE("/deleteTag/:id", v1.DeleteTag)
	}
	//操作文章
	contorllerArticles := rou.Group("/controller/articles")
	//contorllerArticles.Use(middleware.JWT())
	{
		//获取文章列表
		contorllerArticles.GET("/getArticlesList", v1.GetArticlesList)
		//获取指定文章
		contorllerArticles.GET("/getArticles/:id", v1.GetArticle)
		//新建文章
		contorllerArticles.POST("/addArticles", v1.AddArticle)
		//更新指定文章
		contorllerArticles.POST("/editArticles/:id", v1.EditArticle)
		//删除指定文章
		contorllerArticles.DELETE("/deleteArticles/:id", v1.DeleteArticle)
	}

	return rou
}
