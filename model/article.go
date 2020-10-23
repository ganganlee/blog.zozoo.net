package model

//文章表
import (
	"github.com/xormplus/xorm"
	"time"
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