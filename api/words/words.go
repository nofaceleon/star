// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package words

import (
	"context"

	"star2/api/words/v1"
)

type IWordsV1 interface {
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
}
