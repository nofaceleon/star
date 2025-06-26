package middleware

import (
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"star2/internal/consts"
)

// Auth 验证token是否合法
func Auth(r *ghttp.Request) {
	var tokenString = r.Header.Get("Authorization")
	fmt.Println(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(consts.JwtKey), nil

	})
	if err != nil || !token.Valid {
		r.Response.WriteStatus(http.StatusForbidden)
		r.Exit()
	}
	r.Middleware.Next() //放在最后面是前置中间件,在请求之前调用
}
