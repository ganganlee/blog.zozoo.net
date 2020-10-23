package model

import (
	"github.com/xormplus/xorm"
	"time"
)

type (
	Tag struct {
		Id         int64
		Name      string
		Status     int8
		CreateTime time.Time
		UpdateTime time.Time
	}

	TagModel struct {
		sql *xorm.Engine
	}
)
