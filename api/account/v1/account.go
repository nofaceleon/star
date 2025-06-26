package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type InfoReq struct {
	g.Meta `path:"account/info" method:"get" sm:"获取用户信息" tags:"用户"`
}

type InfoRes struct {
	Username  string      `json:"username" dc:"用户名"`
	Email     string      `json:"email" dc:"邮箱"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更行时间"`
}
