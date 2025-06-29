package words

import (
	"context"
	"star2/internal/logic/words"

	"star2/api/words/v1"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {

	uid, err := c.usersLogic.GetUid(ctx)
	if err != nil {
		return nil, err
	}

	err = c.wordsLogic.Update(ctx, req.Id, words.UpdateInput{
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
