package users

import (
	"context"
	"star2/internal/logic/users"

	"star2/api/users/v1"
)


func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	token, err := c.users.Login(ctx, users.LoginInput{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &v1.LoginRes{Token:token}, nil
}
