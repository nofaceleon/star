package users

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v5"
	"star2/internal/consts"
	"star2/internal/dao"
	"star2/internal/model/entity"
	"time"
)

// jwtClaims 我们自己定义的claims
type jwtClaims struct {
	Id       uint
	Username string
	jwt.RegisteredClaims
}

type LoginInput struct {
	Username string
	Password string
}

func (u *Users) Login(ctx context.Context, in LoginInput) (tokenString string, err error) {
	var user entity.Users
	err = dao.Users.Ctx(ctx).Where("username", in.Username).Scan(&user)
	if err != nil {
		return "", gerror.New("用户名或者密码错误")
	}
	if user.Id == 0 {
		return "", gerror.New("用户不存在")
	}

	if user.Password != u.encryptPassword(in.Password) {
		return "", gerror.New("用户名或者密码错误")
	}

	//生成token
	uc := &jwtClaims{
		Id:       user.Id,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(6 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	return token.SignedString([]byte(consts.JwtKey))

}

// Info 获取用户信息
func (u *Users) Info(ctx context.Context) (user *entity.Users, err error) {
	tokenString := g.RequestFromCtx(ctx).Request.Header.Get("Authorization")

	tokenClaims, err := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(consts.JwtKey), nil
	})
	if claims, ok := tokenClaims.Claims.(*jwtClaims); ok && tokenClaims.Valid {
		err = dao.Users.Ctx(ctx).Where("id", claims.Id).Scan(&user)
	}
	return
}

func (u *Users) GetUid(ctx context.Context) (uint, error) {
	user, err := u.Info(ctx)
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}
