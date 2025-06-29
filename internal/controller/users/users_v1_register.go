package users

import (
	"context"
	"star2/internal/logic/users"

	"star2/api/users/v1"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	err = c.usersLogic.Register(ctx, users.RegisterInput{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	})
	return nil, err
}
