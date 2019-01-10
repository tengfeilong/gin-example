package routers

import (
	"gin-example/config"
	"gin-example/middleware"
	"gin-example/routers/contorller/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	rou := gin.New()
	rou.Use(gin.Logger(), gin.Recovery())
	gin.SetMode(config.RunMode)

	rou.GET("/auth", v1.GetAuth)

	//操作标签
	contorllerTag := rou.Group("/controller/tag")
	contorllerTag.Use(middleware.JWT())
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
	contorllerArticles.Use(middleware.JWT())
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
