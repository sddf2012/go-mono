package user

import (
	"context"

	"github.com/sddf2012/go-mono/app/user/api/user/v1"
)

func (c *ControllerV1) GetUser(ctx context.Context, req *v1.GetUserReq) (res *v1.GetUserRes, err error) {
	res = &v1.GetUserRes{Id: 1, Name: "test"}
	return
}
