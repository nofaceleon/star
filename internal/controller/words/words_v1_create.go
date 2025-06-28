package words

import (
	"context"
	"star2/internal/logic/words"

	"star2/api/words/v1"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	uid, err := c.users.GetUid(ctx)
	if err != nil {
		return nil, err
	}

	// 创建单词
	err = c.words.Create(ctx, words.CreateInput{
		Uid:                uid,
		Word:               req.Word,
		Definition:         req.Definition,
		ExampleSentence:    req.ExampleSentence,
		ChineseTranslation: req.ChineseTranslation,
		Pronunciation:      req.Pronunciation,
		ProficiencyLevel:   req.ProficiencyLevel,
	})

	return nil, err
}
