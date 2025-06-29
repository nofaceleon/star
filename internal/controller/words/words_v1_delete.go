package words

import (
	"context"

	"star2/api/words/v1"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {

	uid, err := c.usersLogic.GetUid(ctx)
	if err != nil {
		return nil, err
	}
	err = c.wordsLogic.Delete(ctx, req.Id, uid)
	return nil, err
}
