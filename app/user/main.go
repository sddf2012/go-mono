package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/sddf2012/go-mono/app/user/internal/cmd"
	_ "github.com/sddf2012/go-mono/app/user/internal/packed"
)

func main() {
	go rpcServer()
	cmd.Main.Run(gctx.GetInitCtx())
}

func rpcServer() {
	cmd.GrpcServer.Run(gctx.GetInitCtx())
}
