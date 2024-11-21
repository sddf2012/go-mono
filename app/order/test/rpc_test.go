package test

import (
	"context"
	"flag"
	"fmt"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/net/gsel"
	"github.com/gogf/gf/v2/os/gctx"
	v1 "github.com/sddf2012/go-mono/app/user/api/user/v1/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
	"time"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:18001", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func Test_GetUser(t *testing.T) {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := v1.NewRpcUserServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.RpcGetUser(ctx, &v1.RpcGetUserReq{Id: 1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetName())
}

func Test_GetUser2(t *testing.T) {
	grpcx.Resolver.Register(etcd.New("127.0.0.1:2379"))
	gsel.SetBuilder(gsel.NewBuilderRoundRobin())

	var (
		ctx    = gctx.New()
		conn   = grpcx.Client.MustNewGrpcClientConn("rpc-user")
		client = v1.NewRpcUserServiceClient(conn)
	)

	flag.Parse()
	// Set up a connection to the server.
	r, err := client.RpcGetUser(ctx, &v1.RpcGetUserReq{Id: 1})
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	fmt.Print(r)
}
