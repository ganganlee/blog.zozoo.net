package common

import "github.com/gin-gonic/gin"

//获取gin设置的参数,返回字符串
func GetGinWithString(gin *gin.Context,key string) string {
	val, exists := gin.Get(key)
	if !exists {
		return ""
	}

	return val.(string)
}