package main

import (
	"errors"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	_ "star2/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"star2/internal/cmd"
)

func main() {

	g.I18n().SetLanguage("zh-CN")

	err := connDb()
	if err != nil {
		panic(err)
	}
	cmd.Main.Run(gctx.GetInitCtx())
}

func connDb() error {
	err := g.DB().PingMaster()
	if err != nil {
		return errors.New("数据库连接失败")
	}
	return nil
}
