package user

import (
	"context"
	"github.com/sddf2012/go-mono/app/user/api/user/v1/grpc"
)

func (rpc *RpcControllerV1) RpcGetUser(context.Context, *grpc.RpcGetUserReq) (res *grpc.RpcGetUserRes, err error) {
	res = &grpc.RpcGetUserRes{
		Id:   1,
		Name: "hello",
	}
	return
}
