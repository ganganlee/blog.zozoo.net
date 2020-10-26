package route

import (
	"blog.zozoo.net/common"
	"blog.zozoo.net/config"
	"blog.zozoo.net/controller"
	"blog.zozoo.net/model"
	"github.com/gin-gonic/gin"
	"github.com/xormplus/xorm"
	_ "github.com/go-sql-driver/mysql"
)

//注册路由
func CreateRoute(gin *gin.Engine) {
	//读取配置文件
	conf := config.LoadConfig()

	//注册数据库
	engine, err := xorm.NewEngine(
		"mysql",
		conf.Mysql.Username+":"+conf.Mysql.Password+"@("+conf.Mysql.Host+")/"+conf.Mysql.Database+"?charset=utf8mb4")
	common.HandleErr(err, "连接数据库")


	//注册user控制器
	userModel := model.NewUserModel(engine)
	user := controller.NewUserControl(userModel)

	//注册article控制器
	articleModel := model.NewArticleModel(engine)
	article := controller.NewArticleControl(articleModel)

	gin.GET("/", controller.Index)

	//用户模块
	userGrout := gin.Group("/user")
	userGrout.POST("/register", user.RegisterUser)
	userGrout.POST("/login", user.LoginUser)

	//文章模块
	articleGroup := gin.Group("/article", LoginMiddleWare)
	articleGroup.POST("/create", article.AddArticle)
}
