package model

//文章表
import (
	"github.com/xormplus/xorm"
	"time"
	"errors"
)

type (
	Article struct {
		Id         int64
		UserId     int64
		Title      string
		Content    string
		Sort       int8
		View       int64
		Status     int8
		Comment    int64
		CreateTime time.Time
		UpdateTime time.Time
	}

	ArticleModel struct {
		sql *xorm.Engine
	}
)

func NewArticleModel(sql *xorm.Engine) *ArticleModel  {
	return &ArticleModel{sql: sql}
}

//添加文章
func (a *ArticleModel)CreateArticle(article *Article)error  {
	n, err := a.sql.Insert(article)
	if err != nil {
		return err
	}

	if n == 0 {
		return errors.New("插入失败")
	}

	return nil
}