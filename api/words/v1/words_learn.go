package v1

import "github.com/gogf/gf/v2/frame/g"

type RandListReq struct {
	g.Meta `path:"words/rand" method:"get" sm:"随机单词" tags:"单词"`
	Limit  int `json:"limit" v:"between:1, 300" dc:"限制个数, 默认50" `
}

type RandListRes struct {
	List []SingleWord `json:"list"`
}
