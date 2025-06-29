package words

import (
	"context"
	"star2/internal/logic/words"

	"star2/api/words/v1"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	uid, err := c.usersLogic.GetUid(ctx)
	if err != nil {
		return nil, err
	}
	wordlist, total, err := c.wordsLogic.List(ctx, words.ListInput{
		Uid:  uid,
		Word: req.Word,
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}

	// 包装成需要返回的类型
	var list []v1.SingleWord
	// 遍历数据库返回的数据, 处理成接口需要的数据
	for _, v := range wordlist {
		list = append(list, v1.SingleWord{
			Id:               v.Id,
			Word:             v.Word,
			Definition:       v.Definition,
			ProficiencyLevel: v1.ProficiencyLevel(v.ProficiencyLevel),
		})
	}
	return &v1.ListRes{
		List:  list,
		Total: total,
	}, nil
}
