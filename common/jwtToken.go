package common

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type (
	AccessToken struct {
		Token string
		Secret string
		jwt.StandardClaims
	}
)

//加密key
var jwtSecret = []byte("0^2SljwaYzlRU7*u")

//生成token
func (a *AccessToken)GenerateToken(expireAt int64)  error {

	claims := AccessToken{
		Token: a.Secret,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expireAt,
			// 指定token发行人
			Issuer: "ganganlee",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		return  err
	}

	a.Token = token
	return  nil
}

//验证token
func (a *AccessToken)ValidateToken() (bool,error) {

	token, err := jwt.Parse(a.Token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
		return true,nil
	} else {
		return false,err
	}
}