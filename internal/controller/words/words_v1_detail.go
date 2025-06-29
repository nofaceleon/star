package words

import (
	"context"

	"star2/api/words/v1"
)

func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {

	uid, err := c.usersLogic.GetUid(ctx)
	if err != nil {
		return nil, err
	}
	detail, err := c.wordsLogic.Detail(ctx, req.Id, uid)
	if err != nil {
		return nil, err
	}

	return &v1.DetailRes{
		Id:                 detail.Id,
		Word:               detail.Word,
		Definition:         detail.Definition,
		ExampleSentence:    detail.ExampleSentence,
		ChineseTranslation: detail.ChineseTranslation,
		Pronunciation:      detail.Pronunciation,
		ProficiencyLevel:   v1.ProficiencyLevel(detail.ProficiencyLevel),
		CreatedAt:          detail.CreatedAt,
		UpdatedAt:          detail.UpdatedAt,
	}, nil
}
