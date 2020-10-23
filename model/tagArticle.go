package model

import (
	"github.com/xormplus/xorm"
	"time"
)

type (
	TagArticle struct {
		Id         int64
		TagId      int64
		ArticleId  int64
		UserId     int64
		CreateTime time.Time
		UpdateTime time.Time
	}

	TagArticleModel struct {
		sql *xorm.Engine
	}
)
