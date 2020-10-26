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

	secret, err := accessToken.ValidateToken()
	if err != nil {
		common.ResponseFatal(context,4001,"error",err.Error())
		context.Abort()
	}

	if secret == "" {
		common.ResponseFatal(context,400,"error","token验证失败")
		context.Abort()
	}

	//设置用户secret
	context.Set("userSecret",secret)

	context.Next()
	return
}
