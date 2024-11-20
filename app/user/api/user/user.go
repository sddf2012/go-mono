// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"github.com/sddf2012/go-mono/app/user/api/user/v1"
)

type IUserV1 interface {
	GetUser(ctx context.Context, req *v1.GetUserReq) (res *v1.GetUserRes, err error)
}
