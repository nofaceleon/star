package words

import (
	"context"

	"star2/api/words/v1"
)

func (c *ControllerV1) RandList(ctx context.Context, req *v1.RandListReq) (res *v1.RandListRes, err error) {
	uid, err := c.usersLogic.GetUid(ctx)
	if err != nil {
		return nil, err
	}

	list, err := c.wordsLogic.RandList(ctx, req.Limit, uid)
	if err != nil {
		return nil, err
	}

	var word []v1.SingleWord
	for _, words := range list {
		word = append(word, v1.SingleWord{
			Id:               words.Id,
			Word:             words.Word,
			Definition:       words.Definition,
			ProficiencyLevel: v1.ProficiencyLevel(words.ProficiencyLevel),
		})
	}
	return &v1.RandListRes{
		List: word,
	}, nil
}
