package cmd

import (
	"context"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gsvc"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/sddf2012/go-mono/app/user/internal/controller/user"

	"google.golang.org/grpc"

	"github.com/sddf2012/go-mono/app/user/internal/controller/hello"
)

/*
s := g.Server(`hello.svc`)
*/
var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			gsvc.SetRegistry(etcd.New(`127.0.0.1:2379`))
			s := g.Server("user")
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					hello.NewV1(),
					user.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}

	GrpcServer = gcmd.Command{
		Name:  "grpcServer",
		Usage: "grpcServer",
		Brief: "start grpc server of simple goframe demos",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			grpcx.Resolver.Register(etcd.New("127.0.0.1:2379"))

			c := grpcx.Server.NewConfig()
			c.Options = append(c.Options, []grpc.ServerOption{
				grpcx.Server.ChainUnary(
					grpcx.Server.UnaryValidate,
				)}...,
			)
			s := grpcx.Server.New(c)
			user.Register(s)
			s.Run()
			return nil
		},
	}
)
