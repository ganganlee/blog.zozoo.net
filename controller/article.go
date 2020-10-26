package controller

import (
	"blog.zozoo.net/model"
	"github.com/gin-gonic/gin"
	. "blog.zozoo.net/common"
)

type (
	ArticleControl struct {
		model *model.ArticleModel
	}

	//添加文章结构体
	ArticleRequest struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}
)

func NewArticleControl(model *model.ArticleModel) *ArticleControl  {
	return &ArticleControl{
		model: model,
	}
}

//添加文章
func (a *ArticleControl) AddArticle(c *gin.Context) {
	request := new(ArticleRequest)
	err := c.ShouldBindJSON(request)
	if err != nil {
		ResponseFatal(c, 400, "error", err.Error())
		return
	}

	//article := &model.Article{
	//	UserId: 0,
	//	Title: request.Title,
	//	Content: request.Content,
	//}
}
