package controller

import (
	"blog.zozoo.net/model"
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	. "blog.zozoo.net/common"
	"time"
)

type (

	//用户登陆结构体
	LoginUser struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	//登陆成功响应结构体
	LoginResponse struct {
		Username string `json:"username"`
		Avatar   string `json:"avatar"`
		Token    string `json:"token"`
	}

	//用户注册请求结构体
	RegisterUser struct {
		LoginUser
		Email  string `json:"email" binding:"required"`
		Avatar string `json:"avatar" binding:"required"`
	}

	UserControl struct {
		model *model.UserModel
	}
)

func NewUserControl(model *model.UserModel) *UserControl {
	return &UserControl{
		model: model,
	}
}

//用户注册
func (u *UserControl) RegisterUser(c *gin.Context) {

	//获取请求参数
	registerUser := &RegisterUser{}
	err := c.BindJSON(registerUser)
	if err != nil {
		ResponseFatal(c, 400, err.Error(), "")
		return
	}

	//TODO 参数验证

	//TODO 保存数据
	now := time.Now()

	user := &model.User{
		Username:   registerUser.Username,
		Password:   fmt.Sprintf("%x", md5.Sum([]byte(registerUser.Password))),
		Email:      registerUser.Email,
		Avatar:     registerUser.Avatar,
		Secret:     uuid.NewV4().String(),
		CreateTime: time.Unix(now.Unix(), 0),
		UpdateTime: time.Unix(now.Unix(), 0),
	}

	err = u.model.CreateUser(user)
	if err != nil {
		ResponseFatal(c, 400, err.Error(), "")
		return
	}

	ResponseSuccess(c, 200, "ok", user)
	return
}

//用户登陆
func (u *UserControl) LoginUser(c *gin.Context) {

	//获取请求参数
	login := &LoginUser{}
	err := c.BindJSON(login)
	if err != nil {
		ResponseFatal(c, 400, err.Error(), "")
		return
	}

	//查看数据库判断数据是否存在
	user := &model.User{
		Username: login.Username,
		Password: fmt.Sprintf("%x", md5.Sum([]byte(login.Password))),
	}

	err = u.model.LoginUser(user)
	if err != nil {
		ResponseFatal(c, 400, err.Error(), "")
		return
	}

	//设置token过期时间
	now := time.Now().Add(24 * time.Hour).Unix()

	//获取用户token
	accessToken := new(AccessToken)
	accessToken.Secret = user.Secret
	err = accessToken.GenerateToken(now)
	if err != nil {
		ResponseFatal(c, 400, err.Error(), "")
		return
	}

	//组织返回结构体
	loginResponse := LoginResponse{
		Username: user.Username,
		Avatar:   user.Avatar,
		Token:    accessToken.Token,
	}

	ResponseSuccess(c, 200, "ok", loginResponse)
}

//根据用户secret获取用户信息
func (u *UserControl) GetUserInfoBySecret(secret string) *model.User {
	user := &model.User{Secret: secret}
	err := u.model.GetUserInfoBySecret(user)
	if err != nil {
		return nil
	}

	return user
}
