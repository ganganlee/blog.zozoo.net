package model

//用户表
import (
	"fmt"
	"github.com/xormplus/xorm"
	"time"
	"errors"
)

type (
	User struct {
		Id int64 `json:"id"`
		//用户密钥
		Secret   string `json:"secret" xrom:"varchar(50) notnull unique"`
		Username string `json:"username" xorm:"varchar(125) notnull unique"`
		Password string `json:"password" xorm:"varchar(125) notnull"`
		//用户头像
		Avatar     string    `json:"avatar" xrom:"varchar(125)"`
		Email      string    `json:"email" xorm:"varchar(125)"`
		CreateTime time.Time `json:"create_time" xorm:"DateTime"`
		UpdateTime time.Time `json:"update_time" xorm:"DateTime"`
	}

	UserModel struct {
		sql *xorm.Engine
	}
)

func NewUserModel(sql *xorm.Engine) *UserModel {
	return &UserModel{
		sql: sql,
	}
}

//添加用户
func (u *UserModel) CreateUser(user *User) error {
	n, err := u.sql.Insert(user)
	if err != nil {
		return err
	}

	if n == 0 {
		return fmt.Errorf("添加用户失败")
	}

	return nil
}

//用户登陆
func (u *UserModel) LoginUser(user *User) error {
	ok, err := u.sql.Where("username = ? and password = ?", user.Username, user.Password).Get(user)
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("用户不存在")
	}

	return nil
}

//根据用户私密key获取用户信息
func (u *UserModel)GetUserInfoBySecret(user *User) error  {
	ok, err := u.sql.Where("secret = ?", user.Secret).Get(user)
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("用户不存在")
	}

	return nil
}