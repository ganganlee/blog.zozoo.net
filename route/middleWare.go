package route

import (
	"blog.zozoo.net/common"
	"github.com/gin-gonic/gin"
)

//验证用户登陆中间件
func LoginMiddleWare(context *gin.Context){
	authorization := context.GetHeader("Authorization")

	accessToken := common.AccessToken{
		Token: authorization,
	}
	ok, err := accessToken.ValidateToken()
	if err != nil {
		common.ResponseFatal(context,400,"error",err.Error())
		context.Abort()
	}

	if !ok {
		common.ResponseFatal(context,400,"error","token验证失败")
		context.Abort()
	}

	context.Next()
	return
}
