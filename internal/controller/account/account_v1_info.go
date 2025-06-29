package account

import (
	"context"

	"star2/api/account/v1"
)

func (c *ControllerV1) Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error) {
	// 获取用户信息
	info, err := c.usersLogic.Info(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.InfoRes{
		Username:  info.Username,
		Email:     info.Email,
		CreatedAt: info.CreatedAt,
		UpdatedAt: info.UpdatedAt,
	}, nil
}
