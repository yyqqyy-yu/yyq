package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gomod/dao"
	"gomod/rpc/server/proto"
	"gomod/rpc/server/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"time"
)

func main() {

	dao.Minit("dao/db.yaml")


	s := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: 5 * time.Minute, //这个连接最大的空闲时间，超过就释放，解决proxy等到网络问题（不通知grpc的client和server）
	}))

	proto.RegisterOrderServiceServer(s,new(service.Orderservice))
	lis, err := net.Listen("tcp", "0.0.0.0:7000")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println(lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
